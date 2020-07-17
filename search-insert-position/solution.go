package main

import (
	"sort"

	"fmt"
)

func main() {

	nums := []int{1, 3, 5, 6}
	i := 5
	fmt.Println("Example1:")
	fmt.Println("Expected Output: ", 2)
	fmt.Println("Actual Output: ", searchInsert(nums, i))

	nums = []int{1, 3, 5, 6}
	i = 2
	fmt.Println("Example2:")
	fmt.Println("Expected Output: ", 1)
	fmt.Println("Actual Output: ", searchInsert(nums, i))

	nums = []int{1, 3, 5, 6}
	i = 7
	fmt.Println("Example3:")
	fmt.Println("Expected Output: ", 4)
	fmt.Println("Actual Output: ", searchInsert(nums, i))

	nums = []int{1, 3, 5, 6}
	i = 0
	fmt.Println("Example4:")
	fmt.Println("Expected Output: ", 0)
	fmt.Println("Actual Output: ", searchInsert(nums, i))
}

func searchInsert(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
}
