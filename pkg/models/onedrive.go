package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chandanghosh/azuread-go/config"
)

//OneDrive ..
type OneDrive struct {
	Client *http.Client
}

// GetDriveInfo ..
func (o *OneDrive) GetDriveInfo(authDetails AuthDetails) {
	req, err := http.NewRequest("GET", config.DriveURL, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "bearer "+authDetails.AccessToken)

	res, err := o.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	nbytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("OneDrive Info:", string(nbytes))
}
