package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	exp := [][]int{
		{3},
		{1},
		{2},
		{1, 2, 3},
		{1, 3},
		{2, 3},
		{1, 2},
		{},
	}

	fmt.Println("Example:")
	fmt.Printf("  Input: nums = %v\n", nums)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", subsets(nums))
}

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	backtrack(nums, nil, &result)
	return result
}

func backtrack(nums []int, prev []int, result *[][]int) {
	*result = append(*result, append([]int{}, prev...))
	for i := 0; i < len(nums); i++ {
		backtrack(nums[i+1:], append(prev, nums[i]), result)
	}
}
