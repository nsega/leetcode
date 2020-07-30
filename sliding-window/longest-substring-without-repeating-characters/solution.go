package main

import "fmt"

func main() {
	str := "abcabcbb"
	fmt.Println("Example 1:")
	fmt.Printf("  Input: %s\n", str)
	fmt.Println("  Expected output: ", 3)
	fmt.Println("  Actual output: ", lengthOfLongestSubstring(str))

	str = "bbbbb"
	fmt.Println("Example 2:")
	fmt.Printf("  Input: %s\n", str)
	fmt.Println("  Expected output: ", 1)
	fmt.Println("  Actual output: ", lengthOfLongestSubstring(str))

	str = "pwwkew"
	fmt.Println("Example 3:")
	fmt.Printf("  Input: %s\n", str)
	fmt.Println("  Expected output: ", 3)
	fmt.Println("  Actual output: ", lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	i, j := 0, 0
	max := 0
	set := make(map[byte]bool)
	for j < len(s) {
		if _, ok := set[s[j]]; !ok {
			set[s[j]] = true
			j++
			if len(set) > max {
				max = len(set)
			}

		} else {
			delete(set, s[i])
			i++
		}
	}
	return max
}
