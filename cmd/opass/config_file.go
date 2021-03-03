package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

// AccountCredentials represents 1Password credential data
type AccountCredentials struct {
	signinAddress string
	emailAddress  string
	privateKey    string
}

// SaveCredentials store 1Password sign in credentials to onfig file
func SaveCredentials(signinAddress string, emailAddress string, privateKey string) string {
	ensureConfigFile()
	configFilePath := configFilePath()

	cfg, err := ini.Load(configFilePath)

	if err != nil {
		log.Fatal("Failed to read config file")
	}

	cfg.Section("account").Key("signin_address").SetValue(signinAddress)
	cfg.Section("account").Key("email_address").SetValue(emailAddress)
	cfg.Section("account").Key("secret_key").SetValue(privateKey)

	cfg.SaveTo(configFilePath)
	return configFilePath
}

// PrintCredentials for debug
func PrintCredentials() {
	cfg, err := ini.Load(configFilePath())

	if err != nil {
		fmt.Println("Failed to read config file")
		os.Exit(1)
	}

	fmt.Println("Account signin id:", cfg.Section("account").Key("signin_address").String())
	fmt.Println("Account email address:", cfg.Section("account").Key("email_address").String())
	fmt.Println("Account secret key:", cfg.Section("account").Key("secret_key").String())
}

func ensureConfigFile() {
	if _, err := os.Stat(configFilePath()); os.IsNotExist(err) {
		log.Default().Println("Creating config file at " + configFilePath())
		file, err := os.Create(configFilePath())
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	}
}

func configFilePath() string {
	return os.Getenv("HOME") + "/.opass/config"
}
