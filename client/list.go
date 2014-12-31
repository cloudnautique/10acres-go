package client

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func ListClusterCommand() cli.Command {
	return cli.Command{
		Name:  "list",
		Usage: "List Rancher clusters/nodes",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "n, name",
				Usage: "Name of Cluster to list",
			},
		},
		Action: func(c *cli.Context) {
			listClusterFunc(c)
		},
	}
}

func listClusterFunc(c *cli.Context) {
	fmt.Print("Hello")
}
