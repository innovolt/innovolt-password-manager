package models

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type Group struct {
	Name string
	Id string
}

func (group Group) GetList() []string {
	return []string{
		group.Name,
		group.Id,
	}
}

type Groups struct {
	Items []Group
}

func (groups Groups) IsEmpty() bool {
	return len(groups.Items) == 0
}

func (groups Groups) Render() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Group Name", "Group ID"})
	table.SetRowLine(true)
	blueBoldFgColor := tablewriter.Colors {
		tablewriter.Bold, 
		tablewriter.FgBlueColor,
	}
	table.SetHeaderColor(blueBoldFgColor, blueBoldFgColor)

	greenFgColor := tablewriter.Colors {
		tablewriter.FgGreenColor,
	}
	table.SetColumnColor(greenFgColor, greenFgColor)
	for _, group := range groups.Items {
		table.Append(group.GetList())
	}
	table.Render()
}

