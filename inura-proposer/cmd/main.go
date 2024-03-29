package main

import (
	"os"

	opservice "github.com/inuraorg/inura/inura-service"
	"github.com/urfave/cli/v2"

	"github.com/inuraorg/inura/inura-proposer/flags"
	"github.com/inuraorg/inura/inura-proposer/metrics"
	"github.com/inuraorg/inura/inura-proposer/proposer"
	"github.com/inuraorg/inura/inura-service/cliapp"
	oplog "github.com/inuraorg/inura/inura-service/log"
	"github.com/inuraorg/inura/inura-service/metrics/doc"
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
	app.Name = "inura-proposer"
	app.Usage = "L2Output Submitter"
	app.Description = "Service for generating and submitting L2 Output checkpoints to the L2OutputOracle contract"
	app.Action = curryMain(Version)
	app.Commands = []*cli.Command{
		{
			Name:        "doc",
			Subcommands: doc.NewSubcommands(metrics.NewMetrics("default")),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

// curryMain transforms the proposer.Main function into an app.Action
// This is done to capture the Version of the proposer.
func curryMain(version string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return proposer.Main(version, ctx)
	}
}
