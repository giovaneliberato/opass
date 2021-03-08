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
	tag, itemName := splitTagAndItemName(lookupName)

	if IsTag(tag) {
		if itemName != "" {
			getItemByName(lookupName, itemName, copy)
		} else {
			listItemsByTag(tag)
		}
	} else {
		getItemByName(lookupName, itemName, copy)
	}
}

func ListTags() {
	EnsureAccountSignedIn()
	items := OPGetItems(GetSessionToken())
	itemsByTags := groupItemsByTag(items)
	tags := getTags(itemsByTags)

	root := gotree.New("1Password")
	for _, tag := range tags {
		root.Add(tag)
	}
	fmt.Println(root.Print())

	CacheItems(items)
	CacheTags(itemsByTags)
}

func FlushCachedItems() {
	EnsureAccountSignedIn()
	items := OPGetItems(GetSessionToken())
	itemsByTags := groupItemsByTag(items)

	ClearCache()
	CacheItems(items)
	CacheTags(itemsByTags)
}

func listItemsByTag(tag string) {
	EnsureAccountSignedIn()
	items := OPGetItems(GetSessionToken())
	itemsByTag := groupItemsByTag(items)

	root := gotree.New("1Password")
	tagTree := gotree.New(tag)

	sort.Strings(itemsByTag[tag])
	for _, item := range itemsByTag[tag] {
		tagTree.Add(item)
	}

	root.AddTree(tagTree)
	fmt.Println(root.Print())
}

func getTags(itemsByTags map[string][]string) []string {
	tags := make([]string, 0, len(itemsByTags))
	for t := range itemsByTags {
		tags = append(tags, t)
	}
	sort.Strings(tags)
	return tags
}

func getItemByName(fullName string, name string, copy bool) {
	UUID, err := GetItemUUID(name)
	if err != nil {
		log.Fatalln("Item '" + fullName + "' not found.")
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

func groupItemsByTag(items Items) map[string][]string {
	itemsByTag := make(map[string][]string)
	for _, item := range items {
		if len(item.Overview.Tags) == 0 {
			itemsByTag["untagged"] = append(itemsByTag["untagged"], item.Overview.Title)
		}
		for _, tag := range item.Overview.Tags {
			itemsByTag[tag] = append(itemsByTag[tag], item.Overview.Title)
		}
	}
	return itemsByTag
}

func splitTagAndItemName(name string) (string, string) {
	parts := strings.Split(name, "/")

	if len(parts) == 1 {
		return name, ""
	}

	return parts[0], parts[1]
}
