package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

func OpenIniFile(path string) *ini.File {
	ensureFile(path)

	file, err := ini.Load(path)

	if err != nil {
		log.Fatal("Failed to open file at: " + path)
	}

	return file
}

func ensureFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	}
}
