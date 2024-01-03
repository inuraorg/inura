package op_challenger

import (
	"context"
	"fmt"

	"github.com/inuraorg/inura/inura-challenger/config"
	"github.com/inuraorg/inura/inura-challenger/game"
	"github.com/ethereum/go-ethereum/log"
)

// Main is the programmatic entry-point for running inura-challenger
func Main(ctx context.Context, logger log.Logger, cfg *config.Config) error {
	if err := cfg.Check(); err != nil {
		return err
	}
	service, err := game.NewService(ctx, logger, cfg)
	if err != nil {
		return fmt.Errorf("failed to create the fault service: %w", err)
	}

	return service.MonitorGame(ctx)
}
