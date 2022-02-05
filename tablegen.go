package godiscordutil

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type RawTable [][]string

// Gets maximum width of each column in table for spacing purposes
func getColWidths(table RawTable) (colWidths []string) {
	for i := range table {
		colWidth := 0
		for j := range table[0] {
			if len(table[i][j]) > colWidth {
				colWidth = len(table[i][j])
			}
		}
		colWidths = append(colWidths, strconv.Itoa(colWidth))
	}
	return
}

// Pads cell with spaces to meet width requirements
func padCell(cell string, width int) string {
	if len(cell) == width {
		return cell
	}
	padWidth := width - len(cell)
	return (strings.Repeat(" ", padWidth/2) +
		cell +
		strings.Repeat(" ", (padWidth/2)+(padWidth%2)))
}

// Pads all cells in table so each column has uniform width
func padTable(table RawTable) (spacedTable RawTable) {
	colWidths := getColWidths(table)
	for i := range table {
		for j := range table[0] {
			reqWidth, err := strconv.Atoi(colWidths[i])
			if err != nil {
				log.Println(err)
			}
			table[i][j] = padCell(table[i][j], reqWidth)
		}
	}
	return table
}

// Converts array of strings to table row
func genRow(tableRow []string) string {
	var builder strings.Builder
	builder.WriteString("|")
	for i := range tableRow {
		builder.WriteString(fmt.Sprintf("%s|", tableRow[i]))
	}
	builder.WriteString("\n")
	return builder.String()
}

// Generates simple ascii table as string
func GenTable(data RawTable) string {
	paddedTable := padTable(data)
	var builder strings.Builder

	// Writing header row
	builder.WriteString(genRow(paddedTable[0]))

	// Adding header delineation
	builder.WriteString(
		fmt.Sprintf(
			"%s\n",
			strings.Repeat("-", len(builder.String())-1),
		),
	)

	// Writing body
	for i := range paddedTable {
		builder.WriteString(genRow(paddedTable[i]))
	}
	return builder.String()
}
