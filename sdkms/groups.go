package sdkms

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"innovolt-pm/client"
	"innovolt-pm/common"
)

func GetAllGroups(accountId string) (GetAllGroupsResponse, error) {
	var groups GetAllGroupsResponse
	selectAccountRequest := &SelectAccountRequest{
		Id: accountId,
	}
	err := selectAccount(selectAccountRequest)
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
	request := client.NewRequest()
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
	if resp.StatusCode != http.StatusOK {
		return groups, errors.New("Failed to fetch all the groups. StatusCode: " + resp.Status)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return groups, errors.New("Failed to read response body. Reason: " + err.Error())
	}

	var groupList []GetGroupResponse
	err = json.Unmarshal(bodyText, &groupList)
	if err != nil {
		return groups, err
	}
	groups.Items = groupList

	return groups, nil
}
