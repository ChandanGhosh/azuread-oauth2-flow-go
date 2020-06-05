package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/chandanghosh/azuread-go/config"
	model "github.com/chandanghosh/azuread-go/pkg/models"
)

func main() {

	authorize()

}

func getHTTPClient() *http.Client {
	return http.DefaultClient
}

func authorize() {
	var aURL = config.AuthURL + "?client_id=" + config.ClientID +
		"&scope=Files.ReadWrite Files.ReadWrite.all Sites.Read.All Sites.ReadWrite.All offline_access" +
		"&response_type=code" + "&redirect_uri=" + config.RedirectURL

	authDetails := &model.AuthDetails{Client: getHTTPClient()}
	codeURL, err := authDetails.GetCodeFromPrompt(aURL)
	if err != nil {
		log.Fatalln(err)

	}

	authCode := authDetails.ParseAuthCodeFromURL(codeURL)
	fmt.Printf("code:\n%s", authCode)

	authDetails.GetAccessTokenForCode(authCode)
	userDetails := model.UserDetails{Client: getHTTPClient()}

	userDetails.GetUserInfo(*authDetails)


	driveInfo := &model.OneDrive{Client: getHTTPClient()}
	driveInfo.GetDriveInfo(*authDetails)

}

func getAccessTokenForCode(authCode string) {
	client := getHTTPClient()
	url := "client_id=" + config.ClientID +
		"&grant_type=authorization_code" +
		"&redirect_uri=" + config.RedirectURL +
		"&code=" + authCode

	req, err := http.NewRequest("POST", config.TokenURL, strings.NewReader(url))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	authDetails := &model.AuthDetails{}
	json.NewDecoder(res.Body).Decode(authDetails)
	fmt.Printf("\n\naccess token:%s\n\n", authDetails.AccessToken)
	getUserInfo(*authDetails)
	//getDriveInfo(*authDetails)
}

func getUserInfo(authDetails model.AuthDetails) {
	req, err := http.NewRequest("GET", config.MeURL, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "bearer "+authDetails.AccessToken)

	httpClient := getHTTPClient()
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	var userDetails model.UserDetails
	json.NewDecoder(res.Body).Decode(&userDetails)

	fmt.Println(userDetails.DisplayName)
}

func getDriveInfo(authDetails model.AuthDetails) {
	req, err := http.NewRequest("GET", config.DriveURL, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "bearer "+authDetails.AccessToken)

	httpClient := getHTTPClient()
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	nbytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(nbytes))
}
