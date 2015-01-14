package main

import (
	"github.com/codegangsta/cli"
	"github.com/codeskyblue/go-sh"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}
	app.Action = func(c *cli.Context) {
		name := "someone"
		if len(c.Args()) > 0 {
			name = c.Args()[0]
		}
		if c.String("lang") == "spanish" {
			println("Hola", name)
		} else {
			println("Hello", name)
		}
		sh.Command("echo", "hello").Run()
	}
	app.Commands = []cli.Command{
		{
			Name:      "create",
			ShortName: "c",
			Usage:     "create a new application",
			Action: func(c *cli.Context) {
				println("Creating application")
				if len(c.Args()) > 0 {
					application := c.Args()[0]
					sh.Command("aws", "elasticbeanstalk", "create-application", "--application-name", application).Run()
				}
			},
		},
	}
	app.Run(os.Args)
}
