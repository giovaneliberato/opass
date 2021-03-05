package main

import (
	"github.com/alecthomas/kong"
)

var cli struct {
	LoginName struct {
		LoginName string `arg`
	} `arg hidden`
	AllLogins struct {
	} `cmd hidden default:1`
	Config struct {
	} `cmd help:"Initiate 1Password credentials configuration."`

	Signin struct {
	} `cmd help:"Signin to 1Password using predefined credentials."`

	Vaults struct {
	} `cmd help:"List account vaults."`
}

func main() {
	ctx := kong.Parse(&cli)
	switch ctx.Command() {
	case "<login-name>":
		GetLoginByName(ctx.Args[0])
	case "all-logins":
		ListLogins()
	case "config":
		ConfigAccount()
	case "signin":
		SignInToAccount()
	case "vaults":
		ListVaults()
	}
}
