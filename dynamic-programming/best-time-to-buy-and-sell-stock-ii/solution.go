package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Example 1:")
	fmt.Println("Excepted output: 7")
	fmt.Println("output:", maxProfit(prices))

	prices = []int{1, 2, 3, 4, 5}
	fmt.Println("Example 2:")
	fmt.Println("Excepted output: 4")
	fmt.Println("output:", maxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println("Example 3:")
	fmt.Println("Excepted output: 0")
	fmt.Println("output:", maxProfit(prices))

}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var profit int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
