package sdkms

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"innovolt-pm/client"
	"innovolt-pm/common"
	"innovolt-pm/models"
)

func GetAllAccounts() (models.Accounts, error) {
	var accounts models.Accounts

	accessToken, err := common.GetAccessToken()
	if err != nil {
		return accounts, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", GetAccountsEndpoint, nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)

	if resp.StatusCode == http.StatusUnauthorized {
		return accounts, errors.New("Your session is expired. Please login again.")
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return accounts, errors.New("Failed to read response body. Reason: " + err.Error())
	}

	var accountsList []interface{}
	err = json.Unmarshal(bodyText, &accountsList)
	if err != nil {
		return accounts, err
	}

	for _, account := range accountsList {
		accountDict, _ := account.(map[string]interface{})
		name, _ := accountDict["name"].(string)
		acctId, _ := accountDict["acct_id"].(string)
		account := models.Account{
			Name: name,
			Id:   acctId,
		}
		accounts.Items = append(accounts.Items, account)
	}

	return accounts, nil
}

func SelectAccount(accountId string) error {
	accessToken, err := common.GetAccessToken()
	if err != nil {
		return err
	}

	client := &http.Client{}
	data := map[string]string{"acct_id": accountId}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(data)
	req, err := http.NewRequest("POST", SelectAccountEndpoint, payloadBuf)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New("Your session is expired. Please login again.")
	}
	return nil
}

func GetAllGroups(accountId string) (models.Groups, error) {
	var groups models.Groups
	err := SelectAccount(accountId)
	if err != nil {
		return groups, nil
	}

	accessToken, err := common.GetAccessToken()
	if err != nil {
		return groups, err
	}

	client.SetBearerAuth(accessToken)
	authHeaderVal, err := client.GetAuthHeaderValue()
	if err != nil {
		return groups, err
	}
	request := client.New()
	err = request.WithHeader("Authorization", authHeaderVal)
	if err != nil {
		return groups, err
	}
	err = request.WithMethod("GET")
	if err != nil {
		return groups, err
	}
	err = request.WithUrl(GetGroupsEndpoint)
	if err != nil {
		return groups, err
	}
	resp, err := request.Send()
	// Check for errors because of invalid request etc.
	if err != nil {
		return groups, nil
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return groups, errors.New("Your session is expired. Please login again.")
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return groups, errors.New("Failed to read response body. Reason: " + err.Error())
	}

	var groupsList []interface{}
	err = json.Unmarshal(bodyText, &groupsList)
	if err != nil {
		return groups, err
	}

	for _, group := range groupsList {
		groupDict, _ := group.(map[string]interface{})
		name, _ := groupDict["name"].(string)
		groupId, _ := groupDict["group_id"].(string)
		group := models.Group{
			Name: name,
			Id:   groupId,
		}
		groups.Items = append(groups.Items, group)
	}

	return groups, nil
}

func CreateSecret(secret *models.Secret) error {
	accessToken, err := common.GetAccessToken()
	if err != nil {
		return err
	}

	client.SetBearerAuth(accessToken)
	authHeaderVal, err := client.GetAuthHeaderValue()
	if err != nil {
		return err
	}
	request := client.New()
	err = request.WithHeader("Authorization", authHeaderVal)
	if err != nil {
		return err
	}
	err = request.WithMethod("PUT")
	if err != nil {
		return err
	}
	err = request.WithUrl(PostKeyEndpoint)
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"Name":    secret.Name,
		"GroupId": secret.GroupId,
		"ObjType": "SECRET",
		"KeyOps":  []string{"EXPORT"},
		"Value":   secret.Encode(),
	}
	err = request.WithBody(data)
	if err != nil {
		return nil
	}
	resp, err := request.Send()
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New("Your session is expired. Please login again.")
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Failed to read response body. Reason: " + err.Error())
	}
	// Duplicate/Conflict resource case
	if resp.StatusCode == http.StatusConflict {
		return errors.New("Duplicate secret " + secret.Name)
	}
	if resp.StatusCode != http.StatusCreated {
		return errors.New(string(bodyText))
	}

	return nil
}

func GetSecret(accountId string, groupId string, secretName string) (string, error) {
	accessToken, err := common.GetAccessToken()
	if err != nil {
		return "", err
	}

	client.SetBearerAuth(accessToken)
	authHeaderVal, err := client.GetAuthHeaderValue()
	if err != nil {
		return "", err
	}
	request := client.New()
	err = request.WithHeader("Authorization", authHeaderVal)
	if err != nil {
		return "", err
	}
	err = request.WithMethod("POST")
	if err != nil {
		return "", err
	}
	err = request.WithUrl(ExportKeyEndpoint)
	if err != nil {
		return "", err
	}
	data := map[string]interface{}{
		"name": secretName,
	}
	err = request.WithBody(data)
	if err != nil {
		return "", nil
	}
	resp, err := request.Send()
	if err != nil {
		return "", err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return "", errors.New("Your session is expired. Please login again.")
	}
	if resp.StatusCode == http.StatusNotFound {
		return "", errors.New("Secret is not found in the selected account and group.")
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Failed to read response body. Reason: " + err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(string(bodyText))
	}

	var exportKey map[string]interface{}
	err = json.Unmarshal(bodyText, &exportKey)
	if err != nil {
		return "", err
	}

	value, ok := exportKey["value"].(string)
	if !ok {
		return "", errors.New("Failed to type cast key value into string")
	}

	return value, nil
}
