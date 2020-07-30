package main

import "fmt"

func main() {

	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println("Example 1")
	fmt.Println("Expected output: 3")
	fmt.Println("output:", coinChange(coins, amount))

	coins = []int{2}
	amount = 3
	fmt.Println("Example 2")
	fmt.Println("Expected output: -1")
	fmt.Println("output:", coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min := -1
		for _, coin := range coins {
			if coin > i || dp[i-coin] == -1 {
				continue
			}
			if min == -1 || dp[i-coin]+1 < min {
				min = dp[i-coin] + 1
			}
		}
		dp[i] = min
	}
	return dp[amount]
}
