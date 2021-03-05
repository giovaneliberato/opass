package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

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

func ListTags() {
	EnsureAccountSignedIn()
	items := OPGetLoginItems(GetSessionToken())
	CacheLoginItems(items)

	root := gotree.New("1 Password")
	loginTree := loginsOrderedByTag(items)
	tags := make([]string, 0, len(loginTree))

	for t := range loginTree {
		tags = append(tags, t)
	}
	sort.Strings(tags)

	for _, tag := range tags {
		root.Add(tag)
	}
	fmt.Println(root.Print())
}

func loginsOrderedByTag(items LoginItems) map[string][]string {
	tree := make(map[string][]string)
	for _, item := range items {
		if len(item.Overview.Tags) == 0 {
			tree["untagged"] = append(tree["untagged"], item.Overview.Title)
		}
		for _, tag := range item.Overview.Tags {
			tree[tag] = append(tree[tag], item.Overview.Title)
		}
	}
	return tree
}
