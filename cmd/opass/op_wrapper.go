package main

import (
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

	sessionToken, err := cmd.Output()

	if err != nil {
		os.Exit(1)
	}

	return strings.TrimSuffix(string(sessionToken), "\n")
}

func OPCheckAccountIsSignedIn(sessionToken string) error {
	cmd := exec.Command(
		"op",
		"get",
		"account",
		"--session="+sessionToken)

	_, err := cmd.Output()

	return err
}
