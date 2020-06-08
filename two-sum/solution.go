package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	matched := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == (nums[i] + nums[j]) {
				matched = append(matched, i)
				matched = append(matched, j)
				break
			}
		}
	}
	return matched
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	out := twoSum(nums, target)
	fmt.Println(out)
}
