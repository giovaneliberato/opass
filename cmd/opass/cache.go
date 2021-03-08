package main

import (
	"fmt"
	"log"
	"os"
)

var cacheFileName string = os.Getenv("HOME") + "/.opass/cache"

func GetItemUUID(itemName string) string {
	cache := OpenIniFile(cacheFileName)

	section, err := cache.GetSection("items")

	if err != nil {
		log.Fatalln("Internal cache file is corrupted. Try running 'opass flush'")
	}

	UUID := section.Key(itemName)
	if UUID == nil {
		return ""
	}

	return UUID.String()
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
