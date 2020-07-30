package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{1, 3, 5, 6}
	i := 5
	fmt.Println("searchInsert")
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

	nums = []int{1, 3, 5, 6}
	i = 5
	fmt.Println("====================")
	fmt.Println("searchInsert2")
	fmt.Println("Example1:")
	fmt.Println("Expected Output: ", 2)
	fmt.Println("Actual Output: ", searchInsert2(nums, i))

	nums = []int{1, 3, 5, 6}
	i = 2
	fmt.Println("Example2:")
	fmt.Println("Expected Output: ", 1)
	fmt.Println("Actual Output: ", searchInsert2(nums, i))

	nums = []int{1, 3, 5, 6}
	i = 7
	fmt.Println("Example3:")
	fmt.Println("Expected Output: ", 4)
	fmt.Println("Actual Output: ", searchInsert2(nums, i))

	nums = []int{1, 3, 5, 6}
	i = 0
	fmt.Println("Example4:")
	fmt.Println("Expected Output: ", 0)
	fmt.Println("Actual Output: ", searchInsert2(nums, i))
}

func searchInsert(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
}

func searchInsert2(nums []int, target int) int {
	low := 0
	hi := len(nums) - 1
	mid := hi / 2

	if target > nums[hi] {
		return hi + 1
	}
	for low <= hi {
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			hi = mid - 1
		}
		mid = (hi - low) + low
	}
	return low
}
