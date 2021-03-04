package main

import (
	"fmt"

	"github.com/disiqueira/gotree"
)

func ListVaults() {
	EnsureAccountSignedIn()
}

func ListLogins() {
	EnsureAccountSignedIn()
	items := OPGetLoginItems(GetSessionToken())

	loginItems := gotree.New("logins")

	for _, item := range items {
		loginItems.Add(item.Overview.Title)

	}
	fmt.Println(loginItems.Print())
}
