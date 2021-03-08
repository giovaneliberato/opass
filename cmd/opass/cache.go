package main

import (
	"errors"
	"fmt"
	"os"
)

var cacheFileName string = os.Getenv("HOME") + "/.opass/cache"

func GetItemUUID(itemName string) (string, error) {
	cache := OpenIniFile(cacheFileName)

	section := cache.Section("items")
	UUID := section.Key(itemName).String()

	if UUID == "" {
		return "", errors.New("Item '" + itemName + "' not found")
	}

	return UUID, nil
}

func IsTag(name string) bool {
	cache := OpenIniFile(cacheFileName)
	treeCache := cache.Section("tags")

	return treeCache.HasKey(name)
}

func CacheTags(tags map[string][]string) {
	cache := OpenIniFile(cacheFileName)

	treeCache := cache.Section("tags")

	for tag, items := range tags {
		treeCache.Key(tag).SetValue(fmt.Sprint(len(items)))
	}

	cache.SaveTo(cacheFileName)
}

func CacheItems(items Items) {
	cache := OpenIniFile(cacheFileName)

	treeCache := cache.Section("items")

	for _, item := range items {
		treeCache.Key(item.Overview.Title).SetValue(item.UUID)
	}

	cache.SaveTo(cacheFileName)
}

func ClearCache() {
	cache := OpenIniFile(cacheFileName)
	cache.DeleteSection("items")
	cache.DeleteSection("tags")
	cache.SaveTo(cacheFileName)
}
