package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	exp := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	fmt.Println("Example:")
	fmt.Printf("  Input: nums = %v\n", nums)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", permute(nums))
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	ret := make([][]int, 0)
	backtrack(nums, nil, &ret)
	return ret
}

func backtrack(nums []int, prev []int, ret *[][]int) {
	if len(nums) == 0 {
		*ret = append(*ret, append([]int{}, prev...))
		return
	}
	for i := 0; i < len(nums); i++ {
		backtrack(
			append(append([]int{}, nums[0:i]...), nums[i+1:]...),
			append(prev, nums[i]),
			ret)
	}
}
