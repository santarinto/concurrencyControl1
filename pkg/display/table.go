package display

import (
	"fmt"
	"strings"
)

type TableRow struct {
	Columns []string
}

func PrintTable(headers []string, rows []TableRow) {
	if len(rows) == 0 {
		fmt.Println("No data to display")
		return
	}

	widths := make([]int, len(headers))
	for i, header := range headers {
		widths[i] = len(header)
	}

	for _, row := range rows {
		for i, col := range row.Columns {
			if i < len(widths) && len(col) > widths[i] {
				widths[i] = len(col)
			}
		}
	}

	printBorder(widths)
	printRow(headers, widths, false)
	printBorder(widths)

	for _, row := range rows {
		printRow(row.Columns, widths, true)
	}

	printBorder(widths)
}

func printBorder(widths []int) {
	fmt.Print("+")
	for _, width := range widths {
		fmt.Print(strings.Repeat("-", width+2))
		fmt.Print("+")
	}
	fmt.Println()
}

func printRow(columns []string, widths []int, rightAlignFirst bool) {
	fmt.Print("|")
	for i, col := range columns {
		if i < len(widths) {
			if i == 0 && rightAlignFirst {
				fmt.Printf(" %*s |", widths[i], col)
			} else {
				fmt.Printf(" %-*s |", widths[i], col)
			}
		}
	}
	fmt.Println()
}
