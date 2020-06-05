package main

import (
	"log"
	"net/http"
	"time"

	"github.com/chandanghosh/azuread-go/pkg/models"
)

func getHTTPClient() *http.Client {
	return &http.Client{Timeout: time.Second * 10}
}

func main() {
	authorize()
}

func authorize() {
	chToken := make(chan models.Token)
	authDetails := &models.AuthDetails{}

	authDetails.GetAccessToken(authDetails.GetDeviceCode(), chToken)
	token := <-chToken
	log.Println(token.ExpiresIn)
	log.Println(token.AccessToken)
}
