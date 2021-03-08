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

func GetItem(lookupName string, copy bool) {
	EnsureAccountSignedIn()
	tag, itemName := getTagAndItemName(lookupName)

	if IsTag(tag) {
		if itemName != "" {
			getItemByName(itemName, copy)
		} else {
			listItemsByTag(tag)
		}
	} else {
		getItemByName(itemName, copy)
	}
}

func ListTags() {
	EnsureAccountSignedIn()
	items := OPGetItems(GetSessionToken())

	root := gotree.New("1Password")
	itemsByTags := itemsOrderedByTag(items)
	tags := make([]string, 0, len(itemsByTags))

	for t := range itemsByTags {
		tags = append(tags, t)
	}
	sort.Strings(tags)

	for _, tag := range tags {
		root.Add(tag)
	}
	fmt.Println(root.Print())

	CacheItems(items)
	CacheTags(itemsByTags)
}

func listItemsByTag(tag string) {
	items := OPGetItems(GetSessionToken())

	root := gotree.New("1 Password")
	tagTree := gotree.New(tag)

	itemsByTag := itemsOrderedByTag(items)

	for _, item := range itemsByTag[tag] {
		tagTree.Add(item)
	}

	root.AddTree(tagTree)
	fmt.Println(root.Print())
}

func getItemByName(name string, copy bool) {
	UUID, err := GetItemUUID(name)
	if err != nil {
		log.Fatal("Could not find item " + name)
	}

	loginItem := OPGetItemByUUID(UUID, GetSessionToken())

	if copy {
		clipboard.WriteAll(loginItem.Password)
		fmt.Println("Password copied to clipboard.")
	} else {
		loginEncoded, _ := json.MarshalIndent(loginItem, "", "  ")
		fmt.Println(string(loginEncoded))
	}
}

func itemsOrderedByTag(items Items) map[string][]string {
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

func getTagAndItemName(name string) (string, string) {
	parts := strings.Split(name, "/")

	if len(parts) == 1 {
		return name, ""
	}

	return parts[0], parts[1]
}
