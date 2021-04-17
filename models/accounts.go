package models

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type Account struct {
	Name string
	Id   string
}

func (a Account) GetList() []string {
	return []string{
		a.Name,
		a.Id,
	}
}

type Accounts struct {
	Items []Account
}

func (a Accounts) IsEmpty() bool {
	return len(a.Items) == 0
}

func (a Accounts) Render() {
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
	for _, account := range a.Items {
		table.Append(account.GetList())
	}
	table.Render()
}
