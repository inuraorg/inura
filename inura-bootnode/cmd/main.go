package main

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/inuraorg/inura/inura-bootnode/bootnode"
	"github.com/inuraorg/inura/inura-bootnode/flags"
	oplog "github.com/inuraorg/inura/inura-service/log"
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Name = "bootnode"
	app.Usage = "Rollup Bootnode"
	app.Description = "Broadcasts incoming P2P peers to each other, enabling peer bootstrapping."
	app.Action = bootnode.Main

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
