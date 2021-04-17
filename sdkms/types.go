package sdkms

import (
	"encoding/base64"
	"encoding/json"

	"innovolt-pm/common"
	"innovolt-pm/models"
)

type Secret struct {
	Name     string
	Owner    string
	Domain   string
	Username string
	Password string
}

func (s Secret) ToModel() models.Secret {
	return models.Secret{
		Name:     s.Name,
		Domain:   s.Domain,
		Username: s.Username,
		Password: s.Password,
	}
}

func (s Secret) Render() {
	s.ToModel().Render()
}

func (s Secret) Encode() (string, error) {
	// Serialize the secretValue to JSON bytes
	sBytes, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sBytes), nil
}

func DecodeSecret(encodedSecret string) (Secret, error) {
	decodedSecret, err := base64.StdEncoding.DecodeString(encodedSecret)
	var secret Secret
	if err != nil {
		return secret, nil
	}
	err = json.Unmarshal(decodedSecret, &secret)
	if err != nil {
		return secret, err
	}
	if secret.Owner == common.Owner() {
		return secret, nil
	}
	return Secret{}, common.NotOwnedByUsError
}

type CreateSecretRequest struct {
	AccountId string   `json:"-"`
	Name      string   `json:"name"`
	GroupId   string   `json:"group_id"`
	KeyOps    []string `json:"key_ops"`
	ObjType   string   `json:"obj_type"`
	Value     string   `json:"value"`
}

type SelectAccountRequest struct {
	Id string `json:"acct_id"`
}

type GetAccountResponse struct {
	Id   string `json:"acct_id"`
	Name string `json:"name"`
}

type GetAllAcountsResponse struct {
	Items []GetAccountResponse
}

func (accounts GetAllAcountsResponse) IsEmpty() bool {
	return len(accounts.Items) == 0
}

func (accounts GetAllAcountsResponse) toModel() models.Accounts {
	modelAccounts := models.Accounts{}
	for _, account := range accounts.Items {
		modelAccount := models.Account{
			Id:   account.Id,
			Name: account.Name,
		}
		modelAccounts.Items = append(modelAccounts.Items, modelAccount)
	}
	return modelAccounts
}

func (accounts GetAllAcountsResponse) Render() {
	accounts.toModel().Render()
}

type GetGroupResponse struct {
	Id   string `json:"group_id"`
	Name string `json:"name"`
}

type GetAllGroupsResponse struct {
	Items []GetGroupResponse
}

func (gagr GetAllGroupsResponse) IsEmpty() bool {
	return len(gagr.Items) == 0
}

func (gagr GetAllGroupsResponse) toModel() models.Groups {
	modelGroups := models.Groups{}
	for _, group := range gagr.Items {
		modelGroup := models.Group{
			Id:   group.Id,
			Name: group.Name,
		}
		modelGroups.Items = append(modelGroups.Items, modelGroup)
	}
	return modelGroups
}

func (gagr GetAllGroupsResponse) Render() {
	gagr.toModel().Render()
}

type GetSecretRequest struct {
	Name      string `json:"name"`
	AccountId string `json:"-"`
	GroupId   string `json:"-"`
}

type GetAllSecretsRequest struct {
	AccountId string
	GroupId   string
}
