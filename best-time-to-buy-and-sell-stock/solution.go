package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Example 1:")
	fmt.Println("Excepted output: 5")
	fmt.Println("output:", maxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println("Example 2:")
	fmt.Println("Excepted output: 0")
	fmt.Println("output:", maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var tmp, max int
	for i := 1; i < len(prices); i++ {
		tmp += prices[i] - prices[i-1]
		if tmp < 0 {
			tmp = 0
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}
