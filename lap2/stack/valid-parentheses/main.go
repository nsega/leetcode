package main

import "fmt"

// isValid is Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, ch := range s {
		closer, ok := bracketsMatcher[ch]
		if ok {
			stack = append(stack, closer)
		} else if len(stack) > 0 && ch == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

// bracketsMatcher is rune map the value of which is rune.
var bracketsMatcher = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func main() {
	type args struct {
		testName string
		s        string
		exp      bool
	}
	tests := []args{
		{
			testName: "Example 1",
			s:        "()",
			exp:      true,
		},
		{
			testName: "Example 2",
			s:        "()[]{}",
			exp:      true,
		},
		{
			testName: "Example 3",
			s:        "(]",
			exp:      false,
		},
		{
			testName: "Example 4",
			s:        "([)]",
			exp:      false,
		},
		{
			testName: "Example 5",
			s:        "{[]}",
			exp:      true,
		},
	}
	for _, tt := range tests {
		fmt.Println(tt.testName)
		fmt.Printf("  Input: s = %v\n", tt.s)
		fmt.Println("  Expected:", tt.exp)
		fmt.Println("  Output:", isValid(tt.s))
	}
}
