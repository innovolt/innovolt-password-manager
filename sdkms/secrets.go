package sdkms

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"innovolt-pm/client"
	"innovolt-pm/common"
)

func CreateSecret(createSecretRequest *CreateSecretRequest) error {
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
	err = request.WithMethod("PUT")
	if err != nil {
		return err
	}
	err = request.WithUrl(PostKeyEndpoint)
	if err != nil {
		return err
	}
	requestBytes, err := json.Marshal(createSecretRequest)
	if err != nil {
		return err
	}
	mapInterface, err := common.ToMapInterface(requestBytes)
	if err != nil {
		return err
	}
	err = request.WithBody(mapInterface)
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
		return errors.New("Duplicate secret " + createSecretRequest.Name)
	}
	if resp.StatusCode != http.StatusCreated {
		return errors.New(string(bodyText))
	}

	return nil
}

func GetSecret(getSecretRequest *GetSecretRequest) (Secret, error) {
	var secret Secret
	accessToken, err := common.GetAccessToken()
	if err != nil {
		return secret, err
	}

	client.SetBearerAuth(accessToken)
	authHeaderVal, err := client.GetAuthHeaderValue()
	if err != nil {
		return secret, err
	}
	request := client.NewRequest()
	err = request.WithHeader("Authorization", authHeaderVal)
	if err != nil {
		return secret, err
	}
	err = request.WithMethod("POST")
	if err != nil {
		return secret, err
	}
	err = request.WithUrl(ExportKeyEndpoint)
	if err != nil {
		return secret, err
	}
	requestBytes, err := json.Marshal(getSecretRequest)
	if err != nil {
		return secret, err
	}
	data, err := common.ToMapInterface(requestBytes)
	if err != nil {
		return secret, err
	}
	err = request.WithBody(data)
	if err != nil {
		return secret, nil
	}
	response, err := request.Send()
	if err != nil {
		return secret, err
	}
	if response.StatusCode == http.StatusUnauthorized {
		return secret, errors.New("Your session is expired. Please login again.")
	}
	if response.StatusCode == http.StatusNotFound {
		return secret, errors.New("Secret is not found in the selected account and group.")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return secret, errors.New("Failed to read response body. Reason: " + err.Error())
	}
	if response.StatusCode != http.StatusOK {
		return secret, errors.New(string(body))
	}

	exportedKey, err := common.ToMapInterface(body)
	if err != nil {
		return secret, nil
	}
	value, ok := exportedKey["value"].(string)
	if !ok {
		return secret, errors.New("Failed to type cast key value into string")
	}

	secret, err = DecodeSecret(value)
	if err != nil {
		return secret, err
	}
	return secret, nil
}

func GetAllSecrets(getAllSecretsRequest *GetAllSecretsRequest) ([]Secret, error) {
	var secrets []Secret
	accessToken, err := common.GetAccessToken()
	if err != nil {
		return secrets, err
	}

	client.SetBearerAuth(accessToken)
	authHeaderVal, err := client.GetAuthHeaderValue()
	if err != nil {
		return secrets, err
	}
	request := client.NewRequest()
	err = request.WithHeader("Authorization", authHeaderVal)
	if err != nil {
		return secrets, err
	}
	err = request.WithMethod("GET")
	if err != nil {
		return secrets, err
	}
	err = request.WithUrl(GetSecurityObjectsEndpoint)
	if err != nil {
		return secrets, err
	}
	resp, err := request.Send()
	if err != nil {
		return secrets, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return secrets, errors.New("Your session is expired. Please login again.")
	}
	if resp.StatusCode == http.StatusNotFound {
		return secrets, errors.New("Secret is not found in the selected account and group.")
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return secrets, errors.New("Failed to read response body. Reason: " + err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return secrets, errors.New(string(bodyText))
	}

	var getSecretRequestList []GetSecretRequest
	err = json.Unmarshal(bodyText, &getSecretRequestList)
	if err != nil {
		return secrets, err
	}

	for _, getSecretRequest := range getSecretRequestList {
		getSecretRequest.AccountId = getAllSecretsRequest.AccountId
		getSecretRequest.GroupId = getAllSecretsRequest.GroupId
		secret, _ := GetSecret(&getSecretRequest)
		if secret.Owner == common.Owner() {
			secrets = append(secrets, secret)
		}
	}

	return secrets, nil
}
