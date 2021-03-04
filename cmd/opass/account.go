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

	SaveCredentials(signinAddress, emailAddress, privateKey)
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
