package main

import "fmt"

func main() {

	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println("Example 1")
	fmt.Println("Expected output: true")
	fmt.Println("output:", workBreak(s, wordDict))

	s = "applepenapple"
	wordDict = []string{"apple", "pen"}
	fmt.Println("Example 2")
	fmt.Println("Expected output: true")
	fmt.Println("output:", workBreak(s, wordDict))

	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println("Example 3")
	fmt.Println("Expected output: false")
	fmt.Println("output:", workBreak(s, wordDict))

}

func workBreak(s string, wordDict []string) bool {

	dict := make(map[string]bool)
	for _, ch := range wordDict {
		dict[ch] = true
	}

	memo := make([]bool, len(s)+1)
	memo[0] = true

	for i := 0; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if memo[j] && dict[s[j:i]] {
				memo[i] = true
				break
			}
		}
	}
	return memo[len(s)]
}
