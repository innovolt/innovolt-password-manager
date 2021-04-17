package models

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type Group struct {
	Name string
	Id   string
}

func (g Group) GetList() []string {
	return []string{
		g.Name,
		g.Id,
	}
}

type Groups struct {
	Items []Group
}

func (g Groups) IsEmpty() bool {
	return len(g.Items) == 0
}

func (g Groups) Render() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Group Name", "Group ID"})
	table.SetRowLine(true)
	blueBoldFgColor := tablewriter.Colors{
		tablewriter.Bold,
		tablewriter.FgBlueColor,
	}
	table.SetHeaderColor(blueBoldFgColor, blueBoldFgColor)

	greenFgColor := tablewriter.Colors{
		tablewriter.FgGreenColor,
	}
	table.SetColumnColor(greenFgColor, greenFgColor)
	for _, group := range g.Items {
		table.Append(group.GetList())
	}
	table.Render()
}
