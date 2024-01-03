package main

import (
	"os"

	opservice "github.com/inuraorg/inura/inura-service"
	oplog "github.com/inuraorg/inura/inura-service/log"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	endpointMonitor "github.com/inuraorg/inura/endpoint-monitor"
)

var (
	Version   = ""
	GitCommit = ""
	GitDate   = ""
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = endpointMonitor.CLIFlags("ENDPOINT_MONITOR")
	app.Version = opservice.FormatVersion(Version, GitCommit, GitDate, "")
	app.Name = "endpoint-monitor"
	app.Usage = "Endpoint Monitoring Service"
	app.Description = ""

	app.Action = endpointMonitor.Main(Version)
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
