package cliapp

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"

	"github.com/inuraorg/inura/inura-service/opio"
)

type fakeLifecycle struct {
	startCh, stopCh chan error
	stopped         bool
	selfClose       context.CancelCauseFunc
}

func (f *fakeLifecycle) Start(ctx context.Context) error {
	select {
	case err := <-f.startCh:
		f.stopped = true
		return err
	case <-ctx.Done():
		f.stopped = true
		return ctx.Err()
	}
}

func (f *fakeLifecycle) Stop(ctx context.Context) error {
	select {
	case err := <-f.stopCh:
		f.stopped = true
		return err
	case <-ctx.Done():
		f.stopped = true
		return ctx.Err()
	}
}

func (f *fakeLifecycle) Stopped() bool {
	return f.stopped
}

var _ Lifecycle = (*fakeLifecycle)(nil)

func TestLifecycleCmd(t *testing.T) {

	appSetup := func(t *testing.T, shareApp **fakeLifecycle) (signalCh chan struct{}, initCh, startCh, stopCh, resultCh chan error) {
		signalCh = make(chan struct{})
		initCh = make(chan error)
		startCh = make(chan error)
		stopCh = make(chan error)
		resultCh = make(chan error)

		// mock an application that may fail at different stages of its lifecycle
		mockAppFn := func(ctx *cli.Context, close context.CancelCauseFunc) (Lifecycle, error) {
			select {
			case <-ctx.Context.Done():
				return nil, ctx.Context.Err()
			case err := <-initCh:
				if err != nil {
					return nil, err
				}
			}

			app := &fakeLifecycle{
				startCh:   startCh,
				stopCh:    stopCh,
				stopped:   false,
				selfClose: close,
			}
			if shareApp != nil {
				*shareApp = app
			}
			return app, nil
		}

		// turn our mock app and system signal into a lifecycle-managed command
		actionFn := LifecycleCmd(mockAppFn)

		// try to shut the test down after being locked more than a minute
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

		// puppeteer system signal interrupts by hooking up the test signal channel as "blocker" for the app to use.
		ctx = opio.WithBlocker(ctx, func(ctx context.Context) {
			select {
			case <-ctx.Done():
			case <-signalCh:
			}
		})
		t.Cleanup(cancel)

		// create a fake CLI context to run our command with
		cliCtx := &cli.Context{
			Context: ctx,
			App: &cli.App{
				Name:   "test-app",
				Action: actionFn,
			},
			Command: nil,
		}
		// run the command async, it may block etc. The result will be sent back to the tester.
		go func() {
			result := actionFn(cliCtx)
			require.NoError(t, ctx.Err(), "expecting test context to be alive after end still")
			// collect the result
			resultCh <- result
		}()
		t.Cleanup(func() {
			close(signalCh)
			close(initCh)
			close(startCh)
			close(stopCh)
			close(resultCh)
		})
		return
	}

	t.Run("interrupt int", func(t *testing.T) {
		signalCh, _, _, _, resultCh := appSetup(t, nil)
		signalCh <- struct{}{}
		res := <-resultCh
		require.ErrorIs(t, res, interruptErr)
		require.ErrorContains(t, res, "failed to setup")
	})
	t.Run("failed init", func(t *testing.T) {
		_, initCh, _, _, resultCh := appSetup(t, nil)
		v := errors.New("TEST INIT ERRROR")
		initCh <- v
		res := <-resultCh
		require.ErrorIs(t, res, v)
		require.ErrorContains(t, res, "failed to setup")
	})
	t.Run("interrupt start", func(t *testing.T) {
		var app *fakeLifecycle
		signalCh, initCh, _, _, resultCh := appSetup(t, &app)
		initCh <- nil
		require.False(t, app.Stopped())
		signalCh <- struct{}{}
		res := <-resultCh
		require.ErrorIs(t, res, interruptErr)
		require.ErrorContains(t, res, "failed to start")
		require.True(t, app.Stopped())
	})
	t.Run("failed start", func(t *testing.T) {
		var app *fakeLifecycle
		_, initCh, startCh, _, resultCh := appSetup(t, &app)
		initCh <- nil
		require.False(t, app.Stopped())
		v := errors.New("TEST START ERROR")
		startCh <- v
		res := <-resultCh
		require.ErrorIs(t, res, v)
		require.ErrorContains(t, res, "failed to start")
		require.True(t, app.Stopped())
	})
	t.Run("graceful shutdown", func(t *testing.T) {
		var app *fakeLifecycle
		signalCh, initCh, startCh, stopCh, resultCh := appSetup(t, &app)
		initCh <- nil
		require.False(t, app.Stopped())
		startCh <- nil
		signalCh <- struct{}{} // interrupt, but at an expected time
		stopCh <- nil          // graceful shutdown after interrupt
		require.NoError(t, <-resultCh, nil)
		require.True(t, app.Stopped())
	})
	t.Run("interrupted shutdown", func(t *testing.T) {
		var app *fakeLifecycle
		signalCh, initCh, startCh, _, resultCh := appSetup(t, &app)
		initCh <- nil
		require.False(t, app.Stopped())
		startCh <- nil
		signalCh <- struct{}{} // start graceful shutdown
		signalCh <- struct{}{} // interrupt before the shutdown process is allowed to complete
		res := <-resultCh
		require.ErrorIs(t, res, interruptErr)
		require.ErrorContains(t, res, "failed to stop")
		require.True(t, app.Stopped()) // still fully closes, interrupts only accelerate shutdown where possible.
	})
	t.Run("failed shutdown", func(t *testing.T) {
		var app *fakeLifecycle
		signalCh, initCh, startCh, stopCh, resultCh := appSetup(t, &app)
		initCh <- nil
		require.False(t, app.Stopped())
		startCh <- nil
		signalCh <- struct{}{} // start graceful shutdown
		v := errors.New("TEST STOP ERROR")
		stopCh <- v
		res := <-resultCh
		require.ErrorIs(t, res, v)
		require.ErrorContains(t, res, "failed to stop")
		require.True(t, app.Stopped())
	})
	t.Run("app self-close", func(t *testing.T) {
		var app *fakeLifecycle
		_, initCh, startCh, stopCh, resultCh := appSetup(t, &app)
		initCh <- nil
		require.False(t, app.Stopped())
		startCh <- nil
		v := errors.New("TEST SELF CLOSE ERROR")
		app.selfClose(v)
		stopCh <- nil
		require.NoError(t, <-resultCh, "self-close is not considered an error")
		require.True(t, app.Stopped())
	})
}
