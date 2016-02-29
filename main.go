package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = ("linter")
	app.Usage = "lint a helm chart"
	app.Action = func(c *cli.Context) {
		println("Happy Linting!")
	}

	app.Run(os.Args)

	app.Commands = []cli.Command{
		{
			Name:  "lint",
			Usage: "applies linting to the chart path passed in",
			Action: func(c *cli.Context) {
				path := c.Args().First()
				lint(path)
			},
		},
		{
			Name:  "rules",
			Usage: "options for chart rules",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new chart linting rule",
					Action: func(c *cli.Context) {
						path := c.Args().First()
						addRules(path)
					},
				},
				{
					Name:  "list",
					Usage: "list chart rules",
					Action: func(c *cli.Context) {
						listRules(c)
					},
				},
				{
					Name:    "remove",
					Aliases: []string{"rm"},
					Usage:   "remove a chart rule",
					Action: func(c *cli.Context) {
						path := c.Args().First()
						removeRule(path)
					},
				},
			},
		},
	}

}

func lint(path string) {
	fmt.Println("coming soon")
}

func addRules(path string) {
	fmt.Println("coming soon")
}

func listRules(c *cli.Context) {
	fmt.Println("coming soon")
}

func removeRule(path string) {
	fmt.Println("coming soon")
}
