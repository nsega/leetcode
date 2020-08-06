package main

import (
	"fmt"
	"math"
)

func main() {
	s := "42"
	exp := 42
	fmt.Println("Example 1:")
	fmt.Printf("  Input: s = %v \n", s)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", myAtoi(s))

	s = "   -42"
	exp = -42
	fmt.Println("Example 2:")
	fmt.Printf("  Input: s = %v \n", s)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", myAtoi(s))

	s = "4193 with words"
	exp = 4193
	fmt.Println("Example 3:")
	fmt.Printf("  Input: s = %v \n", s)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", myAtoi(s))

	s = "words and 987"
	exp = 0
	fmt.Println("Example 4:")
	fmt.Printf("  Input: s = %v \n", s)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", myAtoi(s))

	s = "-91283472332"
	exp = -2147483648
	fmt.Println("Example 5:")
	fmt.Printf("  Input: s = %v \n", s)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", myAtoi(s))
}

func myAtoi(str string) int {
	var ans int64 = 0
	var sign int64 = 1
	var start bool = false

	for _, ch := range str {
		if ch <= '9' && ch >= '0' {
			if !start {
				start = true
			}
			ans = ans*10 + int64(ch) - int64('0')
			if ans*sign > math.MaxInt32 {
				return math.MaxInt32
			} else if ans*sign < math.MinInt32 {
				return math.MinInt32
			}
		} else if !start && ch == '+' {
			start = true
		} else if !start && ch == '-' {
			start = true
			sign = -1
		} else if !start && ch == ' ' {
			continue
		} else {
			break
		}
	}
	return int(ans * sign)
}
