package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/disiqueira/gotree"
)

func GetItem(name string, copy bool) {
	EnsureAccountSignedIn()
	tag, loginName := getTagAndLogin(name)

	if IsTag(tag) {
		if loginName != "" {
			getLoginByName(loginName, copy)
		} else {
			listItemsByTag(tag)
		}

	} else {
		getLoginByName(loginName, copy)
	}
}

func ListTags() {
	EnsureAccountSignedIn()
	items := OPGetLoginItems(GetSessionToken())

	root := gotree.New("1Password")
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

	CacheLoginItems(items)
	CacheTags(loginTree)
}

func listItemsByTag(tag string) {
	items := OPGetLoginItems(GetSessionToken())

	root := gotree.New("1 Password")
	tagTree := gotree.New(tag)

	loginsByTag := loginsOrderedByTag(items)

	for _, login := range loginsByTag[tag] {
		tagTree.Add(login)
	}

	root.AddTree(tagTree)
	fmt.Println(root.Print())
}

func getLoginByName(name string, copy bool) {
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

func getTagAndLogin(name string) (string, string) {
	parts := strings.Split(name, "/")

	if len(parts) == 1 {
		return name, ""
	}

	return parts[0], parts[1]
}
