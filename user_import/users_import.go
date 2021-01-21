package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type User struct {
	FirstName     string            `json:"firstName"`
	LastName      string            `json:"lastName"`
	Username      string            `json:"username"`
	Email         string            `json:"email"`
	EmailVerified bool              `json:"emailVerified"`
	Credentials   []Credentials     `json:"credentials"`
	Enabled       bool              `json:"enabled"`
	RealmRoles    []string          `json:"realmRoles"`
	Attributes    map[string]string `json:"attributes"`
}

type Credentials struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Temporary bool   `json:"temporary"`
}

func main() {
	baseURL := flag.String("baseurl", "", "")
	realm := flag.String("realm", "", "")
	username := flag.String("username", "", "")
	password := flag.String("password", "", "")
	clientSecret := flag.String("clientsecret", "", "")
	flag.Parse()

	userURL := "https://taiga.dts-stn.com/media/attachments/8/a/c/a/72251cebecc9911773918f4cb2ba56bfd9d513e55b943522dc02905e74d1/personsampledatav2.json"

	client := http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest(http.MethodGet, userURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "keycloak-user-import")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	jsonMap := []map[string]interface{}{}
	if jsonErr := json.Unmarshal(body, &jsonMap); jsonErr != nil {
		panic(jsonErr)
	}

	accessToken, _, err := KeyCloakLogin(string(*baseURL), "master", "admin-cli", string(*username), string(*password), string(*clientSecret))
	if err != nil {
		panic(err)
	}

	for _, _user := range jsonMap {
		user := User{
			FirstName:     _user["personFirstName"].(string),
			LastName:      _user["personLastName"].(string),
			Email:         _user["personEmailAddress"].(string),
			Attributes:    map[string]string{"guid": _user["guid"].(string)},
			Username:      _user["personFirstName"].(string) + "." + _user["personLastName"].(string),
			EmailVerified: true,
			Enabled:       true,
			RealmRoles:    []string{"offline_access", "uma_authorization"},
			Credentials: []Credentials{
				{
					Type:      "password",
					Value:     "Password1",
					Temporary: false,
				},
			},
		}
		if status, err := KeyCloakCreateUser(user, string(*baseURL), string(*realm), accessToken); err != nil || status == 400 || status == 401 || status == 403 || status > 499 {
			panic(err)
		} else if status == 418 {
			panic("Short and stout")
		}
	}
}

func KeyCloakLogin(baseURL, realm, clientID, username, password, clientSecret string) (accessToken, refreshToken string, err error) {
	client := http.Client{Timeout: time.Second * 30}

	authURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", baseURL, realm)

	data := url.Values{}
	data.Add("username", username)
	data.Add("password", password)
	data.Add("client_id", clientID)
	data.Add("grant_type", "password")
	data.Add("client_secret", clientSecret)

	req, err := http.NewRequest(http.MethodPost, authURL, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	jsonMap := make(map[string]interface{})
	if jsonErr := json.Unmarshal(body, &jsonMap); jsonErr != nil {
		return
	}

	accessToken = jsonMap["access_token"].(string)
	refreshToken = jsonMap["refresh_token"].(string)

	return
}

func KeyCloakCreateUser(user User, baseURL, realm, accessToken string) (int, error) {
	client := http.Client{Timeout: time.Second * 30}

	createURL := fmt.Sprintf("%s/admin/realms/%s/users", baseURL, realm)

	jsonData, err := json.Marshal(user)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, createURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	// if res.StatusCode != http.StatusCreated {
	// 	return fmt.Errorf("Creating new user %s resulted with %v", user.Username, res.Status)
	// }

	if res.Body != nil {
		defer res.Body.Close()
	}

	fmt.Printf("User %s created.\n", user.Username)

	return res.StatusCode, err
}
