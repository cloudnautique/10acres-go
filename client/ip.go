package client

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func IpClusterCommand() cli.Command {
	return cli.Command{
		Name:  "ip",
		Usage: "Get the IP of the server in a Rancher Cluster",
		Action: func(c *cli.Context) {
			fmt.Print("IP\n")
		},
	}
}
