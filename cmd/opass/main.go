package main

import (
	"github.com/alecthomas/kong"
)

var cli struct {
	Config struct {
	} `cmd help:"Initiate 1Password credentials configuration."`
}

func main() {
	ctx := kong.Parse(&cli, kong.UsageOnError())

	switch ctx.Command() {
	case "config":
		ConfigAccount()
	}
}
