package utils

import "github.com/urfave/cli/v2"

var (
	EchoFlag = &cli.StringFlag{
		Name:  "echo",
		Usage: "echo",
	}
)
