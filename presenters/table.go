// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache 2.0

package presenters

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func RenderTable(headings []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headings)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetNoWhiteSpace(true)
	table.SetTablePadding("\t")
	table.AppendBulk(data)
	table.Render()
}
