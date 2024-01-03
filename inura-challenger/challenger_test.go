package op_challenger

import (
	"context"
	"testing"

	"github.com/inuraorg/inura/inura-challenger/config"
	"github.com/inuraorg/inura/inura-service/testlog"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

func TestMainShouldReturnErrorWhenConfigInvalid(t *testing.T) {
	cfg := &config.Config{}
	err := Main(context.Background(), testlog.Logger(t, log.LvlInfo), cfg)
	require.ErrorIs(t, err, cfg.Check())
}
