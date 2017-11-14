package main

import (
	"os"

	"github.com/MottainaiCI/mottainai-agent/cmd"

	"github.com/urfave/cli"
)

const MOTTAINAI_VERSION = "0.0000001"

func main() {
	app := cli.NewApp()
	app.Name = "Mottainai Agent"
	app.Usage = "Task/Job Agent"
	app.Version = MOTTAINAI_VERSION
	app.Commands = []cli.Command{
		cmd.Agent,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
