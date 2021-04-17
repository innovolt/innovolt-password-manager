package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"innovolt-pm/client"
	"innovolt-pm/common"
	"innovolt-pm/sdkms"
)

func Authenticate(credential *Credential) {
	userCredential, ok := credential.User.(UserCredential)
	if ok {
		client.SetBasicUserAuth(userCredential.Username, userCredential.Password)
	} else {
		appCredential, ok := credential.App.(AppCredential)
		if !ok {
			fmt.Println("Neither User nor App credentials are set.")
			return
		}
		client.SetBasicAppAuth(appCredential.ApiKey)
	}

	authHeaderValue, err := client.GetAuthHeaderValue()
	if err != nil {
		fmt.Println("Failed to get Basic Auth Header value. Reason: " + err.Error())
		return
	}
	request := client.NewRequest()
	err = request.WithHeader("Authorization", authHeaderValue)
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
	response, err := request.Send()
	// Check for errors because of invalid request etc.
	if err != nil {
		fmt.Println("Failed to authenticate. Reason: " + err.Error())
		return
	}

	// Check for Unauthorized case
	if response.StatusCode == http.StatusUnauthorized {
		fmt.Println("Invalid Username or Password")
		return
	}

	fmt.Println("Logged in successfully")

	body, err := ioutil.ReadAll(response.Body)

	configFilePath := common.GetAuthConfigFilePath()
	err = common.SaveDataToFile(configFilePath, body)
	if err != nil {
		fmt.Println("Failed to write to " + configFilePath + ". Reason: " + err.Error())
		return
	}
}
