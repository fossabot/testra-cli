package utils

import (
	//"github.com/apcera/termtables"
	"github.com/olekukonko/tablewriter"
	"os"
	"fmt"
)

func PrintTab(headers []string, rows [][]string) {
	fmt.Println()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	table.SetCenterSeparator("")

	table.SetHeader(headers)
	table.AppendBulk(rows)

	table.Render()

	fmt.Println()
}
