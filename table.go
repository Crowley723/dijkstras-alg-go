package main

import (
	"dijkstras-alg-go/model"
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func generateTable(paths [][]model.NodeID) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Destination", "Outgoing Link")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, path := range paths {
		tbl.AddRow(path[1], fmt.Sprintf("(%s,%s)", path[0], path[1]))
	}

	tbl.Print()
}
