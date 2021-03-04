package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strings"
)

type LoginItems []struct {
	UUID      string `json:"uuid"`
	VaultUUID string `json:"vaultUuid"`
	Overview  struct {
		Tags  []string `json:"tags"`
		Title string   `json:"title"`
	} `json:"overview,omitempty"`
}

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

func OPGetLoginItems(sessionToken string) LoginItems {
	cmd := exec.Command(
		"op",
		"list",
		"items",
		"--session="+sessionToken)

	items, err := cmd.Output()

	if err != nil {
		os.Exit(1)
	}

	res := LoginItems{}
	err = json.Unmarshal(items, &res)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return res
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
