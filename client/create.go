package client

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func CreateNewClusterCommand() cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "Create a Rancher Cluster",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "n, name",
				Usage: "Name of Cluster to create",
			},
			cli.StringFlag{
				Name:  "c, count",
				Usage: "Number of nodes",
			},
			cli.StringFlag{
				Name:  "o, os",
				Usage: "OS Image",
			},
			cli.StringFlag{
				Name:  "p, privileged",
				Usage: "Set privileged flag on server, agent or all",
			},
			cli.StringFlag{
				Name:  "server-container",
				Usage: "Specify the Rancher Server container in <repo>/<image>[:version] format",
			},
			cli.StringFlag{
				Name:  "agent-container",
				Usage: "Specify the Rancher Agent container in <repo>/<image>[:version] format",
			},
			cli.StringFlag{
				Name:  "agent-registration-url",
				Usage: "Registration URL for the agent",
			},
		},
		Action: func(c *cli.Context) {
			createClusterFunc(c)
		},
	}
}

func createClusterFunc(c *cli.Context) {
	fmt.Print("creating\n")
}
