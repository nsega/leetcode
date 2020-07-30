package main

import "fmt"

func main() {
	nums := []int{2,3,6,7}
	target := 7
	exp := [][]int{
		{7},
		{2, 2, 3},
	}
	fmt.Println("Example 1:")
	fmt.Printf("  Input: nums = %v, target = %d\n", nums, target)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", combinationSum(nums, target))

	nums = []int{2,3,5}
	target = 8
	exp = [][]int{
		{2,2,2,2},
		{2,3,3},
		{3,5},
	}
	fmt.Println("Example 2:")
	fmt.Printf("  Input: nums = %v, target = %d\n", nums, target)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", combinationSum(nums, target))
}

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	backtrack(candidates, target, 0, nil, &result)
    return result
}

func backtrack(nums []int, target, index int, prev []int, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*result = append(*result, append([]int{}, prev...))
	}
	for i := index; i < len(nums); i++ {
		if target-nums[i] >= 0 {
			backtrack(nums, target-nums[i], i, append(prev, nums[i]), result)
		}
	} 
}
