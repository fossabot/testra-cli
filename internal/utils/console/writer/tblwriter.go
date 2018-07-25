package writer

import (
	//"github.com/apcera/termtables"
	"github.com/olekukonko/tablewriter"
	"github.com/testra-tech/testra-cli/internal/constants/strs"
	"os"
)

func PrintTable(headers []string, rows [][]string) {
	renderTbl := func() {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
		table.SetCenterSeparator(strs.EmptyStr)

		table.SetHeader(headers)
		table.AppendBulk(rows)

		table.Render()
	}
	WrapWithNewLines(renderTbl)
}
