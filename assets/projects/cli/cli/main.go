package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

// START OMIT

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "english",
				Usage:   "Language",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "arguments",
				Action: func(cCtx *cli.Context) error {
					if cCtx.String("lang") == "english" {
						fmt.Printf("Your provided %d arguments!\n", cCtx.Args().Len())
						return nil
					}

					fmt.Printf("Donde esta la biblioteca?")
					return nil
				},
			},

			// MIDDLE OMIT

			{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Options for system files",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "Add a new file to track",
						Action: func(cCtx *cli.Context) error {
							// ...
							return nil
						},
					},
				},
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}

// END OMIT
