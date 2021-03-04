package main

import (
	"fmt"
)

func ConfigAccount() {
	var signinAddress string
	var emailAddress string
	var privateKey string

	fmt.Print("Sign in Address: ")
	fmt.Scanln(&signinAddress)

	fmt.Print("Email Address: ")
	fmt.Scanln(&emailAddress)

	fmt.Print("Private Key: ")
	fmt.Scanln(&privateKey)

	configFilePath := SaveCredentials(signinAddress, emailAddress, privateKey)
	fmt.Printf("Config file created at %s", configFilePath)
}

func SignInToAccount() {
	credentials := GetCredentials()
	sessionToken := OPSignIn(credentials)
	SaveSessionToken(sessionToken)
}

func EnsureAccountSignedIn() {
	sessionToken := GetSessionToken()
	if err := OPCheckAccountIsSignedIn(sessionToken); err != nil {
		SignInToAccount()
	}
}
