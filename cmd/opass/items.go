package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/atotto/clipboard"
	"gopkg.in/yaml.v2"
)

func GetItem(lookupName string, copy bool) {
	EnsureAccountSignedIn()
	tag, itemName := splitTagAndItemName(lookupName)

	if !IsTag(tag) || itemName != "" {
		getItemByName(lookupName, itemName, copy)
	} else {
		listItemsByTag(tag)
	}
}

func ListTags() {
	EnsureAccountSignedIn()
	items := OPGetItems(GetSessionToken())
	itemsByTags := groupItemsByTag(items)
	tags := getTags(itemsByTags)

	PrintSimpleTree(tags)
	CacheItems(items)
	CacheTags(itemsByTags)
}

func ListAllItems() {
	EnsureAccountSignedIn()
	items := OPGetItems(GetSessionToken())
	itemsByTags := groupItemsByTag(items)

	PrintMapTree(itemsByTags)
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

	PrintNestedTree(tag, itemsByTag[tag])
}

func getTags(itemsByTags map[string][]string) []string {
	tags := make([]string, 0, len(itemsByTags))

	for t := range itemsByTags {
		tags = append(tags, t)
	}

	return tags
}

func getItemByName(fullName string, name string, copy bool) {
	UUID := GetItemUUID(name)
	if UUID == "" {
		log.Fatalln("Item '" + fullName + "' not found.")
	}

	loginItem := OPGetItemByUUID(UUID, GetSessionToken())

	if copy {
		clipboard.WriteAll(loginItem.Password)
		fmt.Println("Password copied to clipboard.")
	} else {
		loginEncoded, _ := yaml.Marshal(loginItem)
		fmt.Println()
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
