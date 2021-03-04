package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func OPSignIn(credentials AccountCredentials) string {
	cmd := exec.Command(
		"op",
		"signin",
		credentials.signinAddress,
		credentials.emailAddress,
		credentials.secretKey,
		"--raw")

	cmd.Stdin = os.Stdin

	printPrefix()
	sessionToken, err := cmd.Output()

	if err != nil {
		os.Exit(1)
	}

	return strings.TrimSuffix(string(sessionToken), "\n")
}

func printPrefix() {
	fmt.Print("[1Password CLI] ")
}
