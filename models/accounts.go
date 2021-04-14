package models

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type Account struct {
	Name string
	Id   string
}

func (account Account) GetList() []string {
	return []string{
		account.Name,
		account.Id,
	}
}

type Accounts struct {
	Items []Account
}

func (accounts Accounts) IsEmpty() bool {
	return len(accounts.Items) == 0
}

func (accounts Accounts) Render() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Account Name", "Account ID"})
	table.SetRowLine(true)
	// Set color
	blueBoldFgColor := tablewriter.Colors{
		tablewriter.Bold,
		tablewriter.FgBlueColor,
	}
	table.SetHeaderColor(blueBoldFgColor, blueBoldFgColor)

	greenFgColor := tablewriter.Colors{
		tablewriter.FgGreenColor,
	}
	table.SetColumnColor(greenFgColor, greenFgColor)
	for _, account := range accounts.Items {
		table.Append(account.GetList())
	}
	table.Render()
}
