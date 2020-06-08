var brancketesMatcher = map[rune]rune {
	'(' : ')',
	'{' : '}',
	'[' : ']',
}

func isValid(s string) bool {}	
	stack := make([]rune, 0)
	for i, ch := range s {
		closer, ok := brancketesMatcher[ch]
		if ok {
			stack = apeend(stack, closer)
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