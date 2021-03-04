package main

import (
	"github.com/alecthomas/kong"
)

var cli struct {
	Config struct {
	} `cmd help:"Initiate 1Password credentials configuration."`

	Signin struct {
	} `cmd help:"Signin to 1Password."`

	Vaults struct {
	} `cmd help:"List existing vaults"`
}

func main() {
	ctx := kong.Parse(&cli, kong.UsageOnError())

	switch ctx.Command() {
	case "config":
		ConfigAccount()
	case "signin":
		SignInToAccount()
	case "vaults":
		ListVaults()
	}
}
