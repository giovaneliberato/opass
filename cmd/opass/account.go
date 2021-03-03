package main

import (
	"fmt"
)

func ConfigAccount() {
	var signinAddress string
	var emailAddress string
	var privateKey string

	fmt.Print("Sign in Address: \b")
	fmt.Scanln(&signinAddress)

	fmt.Print("Email Address: \b")
	fmt.Scanln(&emailAddress)

	fmt.Print("Private Key: \b")
	fmt.Scanln(&privateKey)

	configFilePath := SaveCredentials(signinAddress, emailAddress, privateKey)
	fmt.Printf("Config file created at %s", configFilePath)
}
