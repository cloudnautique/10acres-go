package client

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func DeleteClusterCommand() cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "Delete a Rancher Cluster",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "n, name",
				Usage: "Name of Cluster to create",
			},
			cli.BoolFlag{
				Name:  "q, quiet",
				Usage: "Do not prompt for deletion",
			},
		},
		Action: func(c *cli.Context) {
			deleteClusterFunc(c)
		},
	}
}

func deleteClusterFunc(c *cli.Context) {
	fmt.Print("Deleting...\n")
}
