package main

import (
	"os"

	heartbeat "github.com/inuraorg/inura/inura-heartbeat"
	"github.com/inuraorg/inura/inura-heartbeat/flags"
	opservice "github.com/inuraorg/inura/inura-service"
	oplog "github.com/inuraorg/inura/inura-service/log"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

var (
	Version   = ""
	GitCommit = ""
	GitDate   = ""
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = opservice.FormatVersion(Version, GitCommit, GitDate, "")
	app.Name = "inura-heartbeat"
	app.Usage = "Heartbeat recorder"
	app.Description = "Service that records opt-in heartbeats from op nodes"
	app.Action = heartbeat.Main(app.Version)
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
