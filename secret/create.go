package secret

import (
	"errors"
	"fmt"

	"innovolt-pm/common"
	"innovolt-pm/sdkms"
)

func CreateSecret(name string) error {
	_, err := common.GetAccessToken()
	if err != nil {
		return errors.New("Please login using innovolt-pm login")
	}

	// Take input from User
	var domain string
	fmt.Printf("Domain [e.g. www.abc.com]: ")
	fmt.Scanln(&domain)

	var username string
	fmt.Printf("Username: ")
	fmt.Scanln(&username)

	var password string
	fmt.Printf("Password: ")
	fmt.Scanln(&password)

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

	secret := sdkms.Secret{
		Name:     name,
		Owner:    common.Owner(),
		Domain:   domain,
		Username: username,
		Password: password,
	}
	encodedSecret, err := secret.Encode()
	if err != nil {
		return err
	}

	createSecretRequest := sdkms.CreateSecretRequest{
		AccountId: accountId,
		Name:      name,
		GroupId:   groupId,
		KeyOps:    []string{"EXPORT"},
		ObjType:   "SECRET",
		Value:     encodedSecret,
	}

	err = sdkms.CreateSecret(&createSecretRequest)
	if err != nil {
		return err
	}

	fmt.Println("Secret is created successfully.")

	return nil
}
