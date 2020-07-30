package main

import "fmt"

func main() {
	n := 3
	exp := []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}
	fmt.Println("Example :")
	fmt.Printf("  Input: n = %d\n", n)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", generateParenthesis(n))
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	backtrack(&result, "", 0, 0, n)
	return result
}

func backtrack(result *[]string, str string, start, end, max int) {
	if len(str) == max*2 {
		*result = append(*result, str)
		return
	}
	if start < max {
		backtrack(result, str+"(", start+1, end, max)
	}
	if end < start {
		backtrack(result, str+")", start, end+1, max)
	}
}
