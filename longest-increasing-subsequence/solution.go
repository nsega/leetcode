package main

import "fmt"

// lengthOfLIS1 is O(n^2) solution
func lengthOfLIS1(nums []int) int {

	maxLen := 0
	s := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		curMaxLen := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && s[j] >= curMaxLen {
				curMaxLen = s[j] + 1
			}
		}
		s[i] = curMaxLen
	}

	for _, v := range s {
		if v > maxLen {
			maxLen = v
		}
	}
	return maxLen
}

// lengthOfLIS2 is O(n log n) solution
func lengthOfLIS2(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	maxLen := 1
	for i := 0; i < len(nums); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] > max {
				max = dp[j]
			}
		}

		dp[i] = max + 1
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}

func main() {

	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("lengthOfLIS1 - output:", lengthOfLIS1(nums))
	fmt.Println("lengthOfLIS2 - output:", lengthOfLIS2(nums))
}
