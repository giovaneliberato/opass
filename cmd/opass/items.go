package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/disiqueira/gotree"
)

func ListVaults() {
	EnsureAccountSignedIn()
}

func GetLoginByName(name string, copy bool) {
	EnsureAccountSignedIn()

	UUID, err := GetLoginUUID(name)
	if err != nil {
		log.Fatal("Could not find login " + name)
	}

	loginItem := OPGetLogin(UUID, GetSessionToken())

	if copy {
		clipboard.WriteAll(loginItem.Password)
		fmt.Println("Password copied to clipboard.")
	} else {

		loginEncoded, _ := json.MarshalIndent(loginItem, "", "  ")
		fmt.Println(string(loginEncoded))
	}
}

func ListLogins() {
	EnsureAccountSignedIn()
	items := OPGetLoginItems(GetSessionToken())
	CacheLoginItems(items)

	loginItems := gotree.New("logins")

	for _, item := range items {
		loginItems.Add(item.Overview.Title)

	}
	fmt.Println(loginItems.Print())
}
