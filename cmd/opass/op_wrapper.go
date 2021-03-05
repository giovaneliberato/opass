package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type LoginItems []struct {
	UUID      string `json:"uuid"`
	VaultUUID string `json:"vaultUuid"`
	Overview  struct {
		Tags  []string `json:"tags"`
		Title string   `json:"title"`
	} `json:"overview,omitempty"`
}

type loginItemResponse struct {
	Details struct {
		Fields []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"fields"`
	} `json:"details"`
	Overview struct {
		Tags []interface{} `json:"tags"`
		URL  string        `json:"url"`
	} `json:"overview"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	ItemVersion int       `json:"itemVersion"`
}

type LoginItem struct {
	Username    string
	Password    string
	URL         string
	UpdatedAt   time.Time
	ItemVersion int
	Tags        []interface{}
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

func OPGetLogin(loginUUID string, sessionToken string) LoginItem {
	cmd := exec.Command(
		"op",
		"get",
		"item",
		loginUUID,
		"--session="+sessionToken)

	items, err := cmd.Output()

	if err != nil {
		os.Exit(1)
	}

	res := loginItemResponse{}
	err = json.Unmarshal(items, &res)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return convertLoginItem(res)
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

func convertLoginItem(res loginItemResponse) LoginItem {
	userAndPassword := make(map[string]string)

	for _, field := range res.Details.Fields {
		userAndPassword[field.Name] = field.Value
	}

	return LoginItem{
		Username:    userAndPassword["username"],
		Password:    userAndPassword["password"],
		URL:         res.Overview.URL,
		Tags:        res.Overview.Tags,
		ItemVersion: res.ItemVersion,
		UpdatedAt:   res.UpdatedAt,
	}
}
