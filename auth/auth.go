package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"innovolt-pm/client"
	"innovolt-pm/common"
	"innovolt-pm/sdkms"
)

func UserAuthenticate(userCredentials *UserCredentials) {
	var username = userCredentials.Username
	var password = userCredentials.Password

	client.SetBasicUserAuth(username, password)
	authHeaderVal, err := client.GetAuthHeaderValue()
	if err != nil {
		fmt.Println("Failed to get Basic Auth Header value. Reason: " + err.Error())
		return
	}
	request := client.New()
	err = request.WithHeader("Authorization", authHeaderVal)
	if err != nil {
		fmt.Println("Failed to set header. Reason: " + err.Error())
		return
	}
	err = request.WithMethod("POST")
	if err != nil {
		fmt.Println("Failed to set method. Reason: " + err.Error())
		return
	}
	err = request.WithUrl(sdkms.SessionAuthEndpoint)
	if err != nil {
		fmt.Println("Failed to set url. Reason: " + err.Error())
		return
	}
	resp, err := request.Send()
	// Check for errors because of invalid request etc.
	if err != nil {
		fmt.Println("Failed to authenticate. Reason: " + err.Error())
		return
	}

	// Check for Unauthorized case
	if resp.StatusCode == http.StatusUnauthorized {
		fmt.Println("Invalid Username or Password")
		return
	}

	fmt.Println("Logged in successfully")

	bodyText, err := ioutil.ReadAll(resp.Body)

	configFilePath := common.GetAuthConfigFilePath()
	err = common.SaveDataToFile(configFilePath, bodyText)
	if err != nil {
		fmt.Println("Failed to write to " + configFilePath + ". Reason: " + err.Error())
		return
	}
}
