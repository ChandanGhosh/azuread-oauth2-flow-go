package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chandanghosh/azuread-go/config"
)

// UserDetails ..
type UserDetails struct {
	DisplayName string `json:"displayName"`
	GivenName   string `json:"givenName"`
	JobTitle    string `json:"jobTitle"`
	Mail        string `json:"mail"`
	MobilePhone string `json:"mobilePhone"`
	Surname     string `json:"surname"`

	Client *http.Client
}

// GetUserInfo ..
func (u *UserDetails) GetUserInfo(authDetails AuthDetails) {
	req, err := http.NewRequest("GET", config.MeURL, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "bearer "+authDetails.AccessToken)

	res, err := u.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(u)
	fmt.Println("User's display name: ", u.DisplayName)
}
