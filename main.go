package main

import (
	"github.com/cloudnautique/10acres/client"
	"github.com/codegangsta/cli"
	"os"
)

var (
	version string = "0.0.1-dev"
)

func main() {
	app := cli.NewApp()
	app.Name = "10acres"
	app.Usage = "cmd"
	app.Version = version
	app.EnableBashCompletion = true
	app.Author = "Rancher Labs"
	app.Email = ""

	app.Commands = []cli.Command{
		client.CreateNewClusterCommand(),
		client.DeleteClusterCommand(),
		client.ListClusterCommand(),
		client.IpClusterCommand(),
	}

	app.Run(os.Args)
}
