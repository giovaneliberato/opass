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
	secretKey     string
}

// SaveCredentials store 1Password sign in credentials to onfig file
func SaveCredentials(signinAddress string, emailAddress string, privateKey string) string {
	ensureConfigFile()
	cfg := getConfigFile()

	cfg.Section("account").Key("signin_address").SetValue(signinAddress)
	cfg.Section("account").Key("email_address").SetValue(emailAddress)
	cfg.Section("account").Key("secret_key").SetValue(privateKey)

	cfg.SaveTo(configFilePath())
	return configFilePath()
}

// GetCredentials
func GetCredentials() AccountCredentials {
	cfg := getConfigFile()

	return AccountCredentials{
		signinAddress: cfg.Section("account").Key("signin_address").String(),
		emailAddress:  cfg.Section("account").Key("email_address").String(),
		secretKey:     cfg.Section("account").Key("secret_key").String(),
	}
}

func SaveSessionToken(sessionToken string) {
	cfg := getConfigFile()

	cfg.Section("session").Key("token").SetValue(sessionToken)
	cfg.SaveTo(configFilePath())
}

func GetSessionToken() string {
	cfg := getConfigFile()
	return cfg.Section("session").Key("token").String()
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

func getConfigFile() *ini.File {
	configFilePath := configFilePath()

	cfg, err := ini.Load(configFilePath)

	if err != nil {
		log.Fatal("Failed to read config file")
	}
	return cfg
}

func configFilePath() string {
	return os.Getenv("HOME") + "/.opass/config"
}
