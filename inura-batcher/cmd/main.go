package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/inuraorg/inura/inura-batcher/batcher"
	"github.com/inuraorg/inura/inura-batcher/flags"
	"github.com/inuraorg/inura/inura-batcher/metrics"
	opservice "github.com/inuraorg/inura/inura-service"
	"github.com/inuraorg/inura/inura-service/cliapp"
	oplog "github.com/inuraorg/inura/inura-service/log"
	"github.com/inuraorg/inura/inura-service/metrics/doc"
	"github.com/inuraorg/inura/inura-service/opio"
	"github.com/ethereum/go-ethereum/log"
)

var (
	Version   = "v0.10.14"
	GitCommit = ""
	GitDate   = ""
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = cliapp.ProtectFlags(flags.Flags)
	app.Version = opservice.FormatVersion(Version, GitCommit, GitDate, "")
	app.Name = "inura-batcher"
	app.Usage = "Batch Submitter Service"
	app.Description = "Service for generating and submitting L2 tx batches to L1"
	app.Action = cliapp.LifecycleCmd(batcher.Main(Version))
	app.Commands = []*cli.Command{
		{
			Name:        "doc",
			Subcommands: doc.NewSubcommands(metrics.NewMetrics("default")),
		},
	}

	ctx := opio.WithInterruptBlocker(context.Background())
	err := app.RunContext(ctx, os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
