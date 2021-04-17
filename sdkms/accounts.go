package sdkms

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"innovolt-pm/client"
	"innovolt-pm/common"
)

func GetAllAccounts() (GetAllAcountsResponse, error) {
	var accounts GetAllAcountsResponse

	accessToken, err := common.GetAccessToken()
	if err != nil {
		return accounts, err
	}
	client.SetBearerAuth(accessToken)
	authHeaderVal, err := client.GetAuthHeaderValue()
	if err != nil {
		return accounts, err
	}
	request := client.NewRequest()
	err = request.WithHeader("Authorization", authHeaderVal)
	if err != nil {
		return accounts, err
	}
	err = request.WithMethod("GET")
	if err != nil {
		return accounts, err
	}
	err = request.WithUrl(GetAccountsEndpoint)
	if err != nil {
		return accounts, err
	}
	resp, err := request.Send()
	if err != nil {
		return accounts, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return accounts, errors.New("Your session is expired. Please login again.")
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return accounts, errors.New("Failed to read response body. Reason: " + err.Error())
	}

	// Note: go is unable Unmarshal array of json object to models.GetAllAccountsResponse
	// Using []GetAccountResponse which is of type same of models.GetAllAccountsResponse.Items
	// does the job
	var accountsList []GetAccountResponse
	err = json.Unmarshal(bodyText, &accountsList)
	if err != nil {
		return accounts, err
	}
	accounts.Items = accountsList
	return accounts, nil
}

func selectAccount(selectAccountRequest *SelectAccountRequest) error {
	accessToken, err := common.GetAccessToken()
	if err != nil {
		return err
	}
	client.SetBearerAuth(accessToken)
	authHeaderVal, err := client.GetAuthHeaderValue()
	if err != nil {
		return err
	}
	request := client.NewRequest()
	err = request.WithHeader("Authorization", authHeaderVal)
	if err != nil {
		return err
	}
	err = request.WithMethod("POST")
	if err != nil {
		return err
	}
	err = request.WithUrl(SelectAccountEndpoint)
	if err != nil {
		return err
	}

	requestBytes, err := json.Marshal(selectAccountRequest)
	if err != nil {
		return err
	}
	mapInterface, err := common.ToMapInterface(requestBytes)
	if err != nil {
		return err
	}
	err = request.WithBody(mapInterface)
	resp, err := request.Send()
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New("Your session is expired. Please login again.")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("Internal Server Error: Failed to select provided Account.")
	}
	return nil
}
