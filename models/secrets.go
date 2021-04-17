package models

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type Secret struct {
	Name      string
	Domain    string
	Username  string
	Password  string
	AccountId string
	GroupId   string
}

type Secrets struct {
	Items []Secret
}

func (s Secret) GetList() []string {
	return []string{
		s.Name,
		s.Domain,
		s.Username,
		s.Password,
	}
}

func (s Secret) Render() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Domain", "Username", "Password"})
	table.SetRowLine(true)
	// Set color
	blueBoldFgColor := tablewriter.Colors{
		tablewriter.Bold,
		tablewriter.FgBlueColor,
	}
	table.SetHeaderColor(
		blueBoldFgColor,
		blueBoldFgColor,
		blueBoldFgColor,
		blueBoldFgColor,
	)

	greenFgColor := tablewriter.Colors{
		tablewriter.FgGreenColor,
	}
	table.SetColumnColor(
		greenFgColor,
		greenFgColor,
		greenFgColor,
		greenFgColor,
	)
	table.Append(s.GetList())
	table.Render()
}

func (s Secrets) Render() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Domain", "Username", "Password"})
	table.SetRowLine(true)
	// Set color
	blueBoldFgColor := tablewriter.Colors{
		tablewriter.Bold,
		tablewriter.FgBlueColor,
	}
	table.SetHeaderColor(
		blueBoldFgColor,
		blueBoldFgColor,
		blueBoldFgColor,
		blueBoldFgColor,
	)

	greenFgColor := tablewriter.Colors{
		tablewriter.FgGreenColor,
	}
	table.SetColumnColor(
		greenFgColor,
		greenFgColor,
		greenFgColor,
		greenFgColor,
	)

	for _, secret := range s.Items {
		table.Append(secret.GetList())
	}

	table.Render()
}
