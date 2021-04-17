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

	getSecretRequest := sdkms.GetSecretRequest{
		Name:      secretName,
		AccountId: accountId,
		GroupId:   groupId,
	}
	secret, err := sdkms.GetSecret(&getSecretRequest)
	if err != nil {
		return err
	}

	secret.Render()

	return nil
}

func GetAllSecret() error {
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

	getAllSecretsRequest := &sdkms.GetAllSecretsRequest{
		AccountId: accountId,
		GroupId:   groupId,
	}
	secrets, err := sdkms.GetAllSecrets(getAllSecretsRequest)
	if err != nil {
		return err
	}

	modelSecrets := models.Secrets{}
	for _, secret := range secrets {
		modelSecrets.Items = append(modelSecrets.Items, secret.ToModel())
	}

	modelSecrets.Render()
	return nil
}
