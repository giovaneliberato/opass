package main

import (
	"github.com/alecthomas/kong"
)

var cli struct {
	Copy       bool `short:c help:"Copy password to clipboard."`
	ListAll    bool `short:a help:"List all tags and."`
	TagOrLogin struct {
		TagOrLogin string `arg`
	} `arg help:"If a Tag name is given, list all logins. If a Login is given, show details."`
	List   struct{} `cmd default:1 help:"List all tags of account."`
	Config struct{} `cmd help:"Initiate 1Password credentials configuration."`
	Signin struct{} `cmd help:"Signin to 1Password using predefined credentials."`
	Flush  struct{} `cmd help:"Drop local list of items and sync with 1Password account."`
}

func main() {
	ctx := kong.Parse(&cli)
	switch ctx.Command() {
	case "<tag-or-login>":
		GetItem(cli.TagOrLogin.TagOrLogin, cli.Copy)
	case "list":
		if cli.ListAll {
			ListAllItems()
		} else {
			ListTags()
		}
	case "config":
		ConfigAccount()
	case "signin":
		SignInToAccount()
	case "flush":
		FlushCachedItems()
	}
}
