package main

import (
	"fmt"
	"os"
)

var configFileName string = os.Getenv("HOME") + "/.opass/config"

// AccountCredentials represents 1Password credential data
type AccountCredentials struct {
	signinAddress string
	emailAddress  string
	secretKey     string
}

// SaveCredentials store 1Password sign in credentials to onfig file
func SaveCredentials(signinAddress string, emailAddress string, privateKey string) {
	configFile := OpenIniFile(configFileName)

	configFile.Section("account").Key("signin_address").SetValue(signinAddress)
	configFile.Section("account").Key("email_address").SetValue(emailAddress)
	configFile.Section("account").Key("secret_key").SetValue(privateKey)

	configFile.SaveTo(configFileName)
	fmt.Printf("Configuration file created at %s", configFileName)
}

// GetCredentials
func GetCredentials() AccountCredentials {
	configFile := OpenIniFile(configFileName)

	return AccountCredentials{
		signinAddress: configFile.Section("account").Key("signin_address").String(),
		emailAddress:  configFile.Section("account").Key("email_address").String(),
		secretKey:     configFile.Section("account").Key("secret_key").String(),
	}
}

func SaveSessionToken(sessionToken string) {
	configFile := OpenIniFile(configFileName)

	configFile.Section("session").Key("token").SetValue(sessionToken)
	configFile.SaveTo(configFileName)
}

func GetSessionToken() string {
	configFile := OpenIniFile(configFileName)
	return configFile.Section("session").Key("token").String()
}
