package models

import (
	"encoding/base64"
	"errors"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Secret struct {
	Name string
	Domain string
	Username string
	Password string
	AccountId string
	GroupId string
}

func (secret Secret) GetList() []string {
	return []string{
		secret.Name,
		secret.Domain,
		secret.Username,
		secret.Password,
	}
}

func (secret Secret) Render() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Domain", "Username", "Password"})
	table.SetRowLine(true)
	// Set color
	blueBoldFgColor := tablewriter.Colors {
		tablewriter.Bold, 
		tablewriter.FgBlueColor,
	}
	table.SetHeaderColor(
		blueBoldFgColor,
		blueBoldFgColor,
		blueBoldFgColor,
		blueBoldFgColor,
	)

	greenFgColor := tablewriter.Colors {
		tablewriter.FgGreenColor,
	}
	table.SetColumnColor(
		greenFgColor,
		greenFgColor,
		greenFgColor,
		greenFgColor,
	)
	table.Append(secret.GetList())
	table.Render()
}

func (secret Secret) Encode() string {
	val := secret.Domain + "," + secret.Username + "," + secret.Password
	return base64.StdEncoding.EncodeToString([]byte(val))
}

func DecodeSecret(secret string) (Secret, error) {
	b64decodedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return Secret{}, errors.New("Invalid base64 secret value.")
	}
	decodedStr := strings.Split(string(b64decodedSecret), ",")
	if len(decodedStr) != 3 {
		return Secret{}, errors.New("Invalid secret found. Unable to decode it.")
	}
	domain := decodedStr[0]
	username := decodedStr[1]
	password := decodedStr[2]

	return Secret {
		Name: "",
		Domain: domain,
		Username: username,
		Password: password,
		AccountId: "",
		GroupId: "",
	}, nil
}
