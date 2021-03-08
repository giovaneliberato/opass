package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Items []struct {
	UUID     string `json:"uuid"`
	Overview struct {
		Tags  []string `json:"tags"`
		Title string   `json:"title"`
	} `json:"overview,omitempty"`
}

type Item struct {
	Username    string
	Password    string
	URL         string
	UpdatedAt   time.Time
	ItemVersion int
	Tags        []interface{}
}

type itemResponse struct {
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

func OPSignIn(credentials AccountCredentials) string {
	cmd := exec.Command(
		"op",
		"signin",
		credentials.signinAddress,
		credentials.emailAddress,
		credentials.secretKey,
		"--raw")

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	sessionToken, err := cmd.Output()

	if err != nil {
		os.Exit(1)
	}

	return strings.TrimSuffix(string(sessionToken), "\n")
}

func OPGetItemByUUID(itemUUID string, sessionToken string) Item {
	cmd := exec.Command(
		"op",
		"get",
		"item",
		itemUUID,
		"--session="+sessionToken)

	items, err := cmd.Output()

	if err != nil {
		os.Exit(1)
	}

	res := itemResponse{}
	err = json.Unmarshal(items, &res)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return convertResponseToItem(res)
}

func OPGetItems(sessionToken string) Items {
	cmd := exec.Command(
		"op",
		"list",
		"items",
		"--session="+sessionToken,
		"--cache")

	items, err := cmd.Output()

	if err != nil {
		os.Exit(1)
	}

	res := Items{}
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

func convertResponseToItem(res itemResponse) Item {
	userAndPassword := make(map[string]string)

	for _, field := range res.Details.Fields {
		userAndPassword[field.Name] = field.Value
	}

	return Item{
		Username:    userAndPassword["username"],
		Password:    userAndPassword["password"],
		URL:         res.Overview.URL,
		Tags:        res.Overview.Tags,
		ItemVersion: res.ItemVersion,
		UpdatedAt:   res.UpdatedAt,
	}
}
