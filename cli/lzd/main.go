package main

import (
	"os"
	"github.com/codegangsta/cli"
	"ria/cli/commands"
	"runtime"
)

func main() {
	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)
	app := cli.NewApp()
	app.Name = "lzd"
	app.Usage = "Distributed lazada log watcher!"
	app.Commands = []cli.Command{
		commands.Tail(),
	}

	app.Run(os.Args)
}

