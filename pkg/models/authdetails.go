package models

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/chandanghosh/azuread-go/config"
)

// AuthDetails ...
type AuthDetails struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	ExtExpiresIn int    `json:"ext_expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`

	Client *http.Client
}

// GetDeviceCode ..
func (a *AuthDetails) GetDeviceCode() Code {
	res, err := http.Post(config.DeviceCodeURL, "application/x-www-form-urlencoded", strings.NewReader("client_id="+config.ClientID+config.Scopes))
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	var c Code
	json.NewDecoder(res.Body).Decode(&c)
	log.Println(c.Message)

	return c
}

// GetAccessToken ..
func (a *AuthDetails) GetAccessToken(code Code, chAccessToken chan<- Token) {

	ticker := time.NewTicker(time.Duration(code.Interval) * time.Second).C
	go func() {
		for {
			select {
			case <-ticker:
				fmt.Printf("Polling the endpoint for access_token every %vs interval\n", code.Interval)
				res, err := http.Post(config.TokenURL, "application/x-www-form-urlencoded", strings.NewReader("client_id="+config.ClientID+config.GrantType+"&code="+code.DeviceCode))
				if err != nil {
					log.Println(err)
				}
				defer res.Body.Close()
				var token Token
				err = json.NewDecoder(res.Body).Decode(&token)
				if err != nil {
					log.Println(err)
				}

				if token.AccessToken != "" {
					fmt.Println(token.AccessToken)
					chAccessToken <- token
					return
				}

				log.Println("Access token not received yet!")
			}
		}
	}()
}

// GetCodeFromPrompt ..
func (a *AuthDetails) GetCodeFromPrompt(fullAuthURL string) (string, error) {
	var code string
	fmt.Printf("Go to the following link in your browser. After completing "+
		"the authorization flow, enter the authorization code on the command "+
		"line: \n\n%v\n\n\n", fullAuthURL)

	fmt.Println("Enter the returned uri from the browser:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		code = scanner.Text()
		break
	}

	return code, nil
}

// ParseAuthCodeFromURL ...
func (a *AuthDetails) ParseAuthCodeFromURL(codeURL string) string {
	u, err := url.Parse(codeURL)
	if err != nil {
		log.Fatalln("Error parsing the uri", err)
	}
	queryValues, _ := url.ParseQuery(u.RawQuery)
	return queryValues.Get("code")
}

// GetAccessTokenForCode ...
func (a *AuthDetails) GetAccessTokenForCode(authCode string) {

	url := "client_id=" + config.ClientID +
		"&grant_type=authorization_code" +
		"&redirect_uri=" + config.RedirectURL +
		"&code=" + authCode

	req, err := http.NewRequest("POST", config.TokenURL, strings.NewReader(url))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := a.Client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(a)
	fmt.Printf("\n\naccess token:%s\n\n", a.AccessToken)
}
