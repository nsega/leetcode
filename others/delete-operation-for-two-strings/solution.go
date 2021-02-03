package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		return len(word1) + len(word2)
	}
	dp := make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		dp[i] = make([]int, len(word2))
	}
	flag := false
	for i := 0; i < len(word1); i++ {
		if word1[i] == word2[0] || flag {
			dp[i][0] = 1
			flag = true
		}
	}
	flag = false
	for i := 0; i < len(word2); i++ {
		if word2[i] == word1[0] || flag {
			dp[0][i] = 1
			flag = true
		}
	}
	for i := 1; i < len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if word1[i] == word2[j] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return len(word1) + len(word2) - 2*dp[len(word1)-1][len(word2)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	input := []string{"", ""}
	expected := 0
	fmt.Println("input", input)
	fmt.Println("expected", expected)
	fmt.Println("actual", minDistance(input[0], input[1]))

	input = []string{"", "hit"}
	expected = 3
	fmt.Println("input", input)
	fmt.Println("expected", expected)
	fmt.Println("actual", minDistance(input[0], input[1]))

	input = []string{"neat", ""}
	expected = 4
	fmt.Println("input", input)
	fmt.Println("expected", expected)
	fmt.Println("actual", minDistance(input[0], input[1]))

	input = []string{"heat", "hit"}
	expected = 2
	fmt.Println("input", input)
	fmt.Println("expected", expected)
	fmt.Println("actual", minDistance(input[0], input[1]))

	input = []string{"hot", "not"}
	expected = 2
	fmt.Println("input", input)
	fmt.Println("expected", expected)
	fmt.Println("actual", minDistance(input[0], input[1]))

	input = []string{"some", "thing"}
	expected = 9
	fmt.Println("input", input)
	fmt.Println("expected", expected)
	fmt.Println("actual", minDistance(input[0], input[1]))

	input = []string{"abc", "adbc"}
	expected = 1
	fmt.Println("input", input)
	fmt.Println("expected", expected)
	fmt.Println("actual", minDistance(input[0], input[1]))

	input = []string{"awesome", "awesome"}
	expected = 0
	fmt.Println("input", input)
	fmt.Println("expected", expected)
	fmt.Println("actual", minDistance(input[0], input[1]))

	input = []string{"ab", "ba"}
	expected = 2
	fmt.Println("input", input)
	fmt.Println("expected", expected)
	fmt.Println("actual", minDistance(input[0], input[1]))
}
