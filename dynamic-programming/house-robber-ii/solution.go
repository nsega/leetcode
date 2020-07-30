package main

import "fmt"

func main() {
	nums := []int{2, 3, 2}
	fmt.Println("Example 1")
	fmt.Println("Excepted output: 3")
	fmt.Println("output: ", rob(nums))

	nums = []int{1, 2, 3, 1}
	fmt.Println("Example 2")
	fmt.Println("Excepted output: 4")
	fmt.Println("output: ", rob(nums))
}

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(helper(nums[:len(nums)-1]), helper(nums[1:]))
}

func helper(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
