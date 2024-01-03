package main

import (
	"context"
	"os"

	op_challenger "github.com/inuraorg/inura/inura-challenger"
	opservice "github.com/inuraorg/inura/inura-service"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/inuraorg/inura/inura-challenger/config"
	"github.com/inuraorg/inura/inura-challenger/flags"
	"github.com/inuraorg/inura/inura-challenger/version"
	"github.com/inuraorg/inura/inura-service/cliapp"
	oplog "github.com/inuraorg/inura/inura-service/log"
)

var (
	GitCommit = ""
	GitDate   = ""
)

// VersionWithMeta holds the textual version string including the metadata.
var VersionWithMeta = opservice.FormatVersion(version.Version, GitCommit, GitDate, version.Meta)

func main() {
	args := os.Args
	if err := run(args, op_challenger.Main); err != nil {
		log.Crit("Application failed", "err", err)
	}
}

type ConfigAction func(ctx context.Context, log log.Logger, config *config.Config) error

func run(args []string, action ConfigAction) error {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Version = VersionWithMeta
	app.Flags = cliapp.ProtectFlags(flags.Flags)
	app.Name = "inura-challenger"
	app.Usage = "Challenge outputs"
	app.Description = "Ensures that on chain outputs are correct."
	app.Action = func(ctx *cli.Context) error {
		logger, err := setupLogging(ctx)
		if err != nil {
			return err
		}
		logger.Info("Starting inura-challenger", "version", VersionWithMeta)

		cfg, err := flags.NewConfigFromCLI(ctx)
		if err != nil {
			return err
		}
		return action(ctx.Context, logger, cfg)
	}
	return app.Run(args)
}

func setupLogging(ctx *cli.Context) (log.Logger, error) {
	logCfg := oplog.ReadCLIConfig(ctx)
	logger := oplog.NewLogger(oplog.AppOut(ctx), logCfg)
	oplog.SetGlobalLogHandler(logger.GetHandler())
	return logger, nil
}
