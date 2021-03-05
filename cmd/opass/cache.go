package main

import (
	"errors"
	"fmt"
	"os"
)

var cacheFileName string = os.Getenv("HOME") + "/.opass/cache"

func GetLoginUUID(itemName string) (string, error) {
	cache := OpenIniFile(cacheFileName)

	UUID := cache.Section("login-items").Key(itemName).String()

	if UUID == "" {
		return UUID, errors.New("Key do not exist")
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

func CacheLoginItems(LoginItems LoginItems) {
	cache := OpenIniFile(cacheFileName)

	treeCache := cache.Section("login-items")

	for _, item := range LoginItems {
		treeCache.Key(item.Overview.Title).SetValue(item.UUID)
	}

	cache.SaveTo(cacheFileName)
}
