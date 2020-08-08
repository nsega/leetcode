package main

import "fmt"

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	exp := "PAHNAPLSIIGYIR"
	fmt.Println("Example 1:")
	fmt.Printf("  Input: s = %v, numRows = %d\n", s, numRows)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output:   ", convert(s, numRows))

	s = "PAYPALISHIRING"
	numRows = 4
	exp = "PINALSIGYAHRPI"
	fmt.Println("Example 2:")
	fmt.Printf("  Input: s = %v, numRows = %d\n", s, numRows)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output:   ", convert(s, numRows))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	runes := []rune(s)
	runesOut := make([]rune, n)
	j := 0
	for row := 0; row < numRows; row++ {
		k := row
		skip := 2 * (numRows - row - 1)
		for j < n && k < n {
			if skip > 0 {
				runesOut[j] = runes[k]
				k = k + skip
				j++
			}
			skip = 2*(numRows-1) - skip
		}
	}
	return string(runesOut)
}
