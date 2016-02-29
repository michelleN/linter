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
	}

}

func lint(path string) {
	fmt.Println("coming soon")
}
