package main

import (
	"gcli/command"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	flags := []cli.Flag{}

	commands := []*cli.Command{
		command.NewServer(),
		command.NewInfo(),
	}

	app := &cli.App{
		Name:  "shcli",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			return nil
		},
		Flags:    flags,
		Commands: commands,
	}

	// 对flag排序
	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// rsrc -arch amd64 -manifest gcli.manifest -ico gcli.ico -o gcli.syso
// go build
