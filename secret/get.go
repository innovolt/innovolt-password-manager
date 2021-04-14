package secret

import (
	"errors"
	"fmt"

	"innovolt-pm/common"
	"innovolt-pm/models"
	"innovolt-pm/sdkms"
)

func GetSecret(secretName string) error {
	_, err := common.GetAccessToken()
	if err != nil {
		return errors.New("Please login using innovolt-pm login")
	}

	accounts, err := sdkms.GetAllAccounts()
	if err != nil {
		return err
	}
	if accounts.IsEmpty() {
		return errors.New("No account is found in SDKMS. Please create one.")
	}

	accounts.Render()

	var accountId string
	fmt.Printf("Select an Account [ID]: ")
	fmt.Scanln(&accountId)

	groups, err := sdkms.GetAllGroups(accountId)
	if err != nil {
		return err
	}
	if groups.IsEmpty() {
		return errors.New("No group is found in SDKMS. Please create one.")
	}
	groups.Render()

	var groupId string
	fmt.Printf("Select a Group [ID]: ")
	fmt.Scanln(&groupId)

	secretVal, err := sdkms.GetSecret(accountId, groupId, secretName)
	if err != nil {
		return err
	}

	secret, err := models.DecodeSecret(secretVal)
	if err != nil {
		return err
	}

	secret.Name = secretName
	secret.AccountId = accountId
	secret.GroupId = groupId
	secret.Render()

	return nil
}
