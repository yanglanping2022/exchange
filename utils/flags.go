package utils

import "github.com/urfave/cli/v2"

var (
	ConfigFlag = &cli.StringFlag{
		Name:  "config",
		Value: "./config.toml",
		Usage: "config xxx.toml",
	}
)
