package main

var bracketsMatcher = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

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
