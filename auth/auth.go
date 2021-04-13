package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"innovolt-pm/common"
	"innovolt-pm/sdkms"
)

func UserAuthenticate(userCredentials *UserCredentials) {
	var username = userCredentials.Username
	var password = userCredentials.Password

    client := &http.Client{}
    req, err := http.NewRequest("POST", sdkms.SessionAuthEndpoint, nil)
    req.SetBasicAuth(username, password)
    resp, err := client.Do(req)
    
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

   	fmt.Println("Logged in successfully" ) 

	bodyText, err := ioutil.ReadAll(resp.Body)

    configFilePath := common.GetAuthConfigFilePath()
	err = common.SaveDataToFile(configFilePath, bodyText)
	if err != nil {
		fmt.Println("Failed to write to " + configFilePath + ". Reason: " + err.Error())
		return
	}
}
