package main

import "fmt"

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i+1] = max(dp[i], dp[i-1]+nums[i])
	}
	return dp[len(nums)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println("Expected output: 4")
	fmt.Println("output:", rob(nums))

	nums = []int{2, 7, 9, 3, 1}
	fmt.Println("Expected output: 12")
	fmt.Println("output:", rob(nums))

}
