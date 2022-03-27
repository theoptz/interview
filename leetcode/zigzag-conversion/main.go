package main

import (
	"fmt"
	"strings"
)

// Follow link to see details
// https://leetcode.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, numRows)
	step := 1
	row := 0

	for i := 0; i < len(s); i++ {
		rows[row] += string(s[i])

		if row == 0 {
			step = 1
		} else if row == numRows-1 {
			step = -1
		}

		row += step
	}

	return strings.Join(rows, "")
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3)) // PAHNAPLSIIGYIR
	fmt.Println(convert("PAYPALISHIRING", 4)) // PINALSIGYAHRPI
	fmt.Println(convert("A", 1))              // A
}
