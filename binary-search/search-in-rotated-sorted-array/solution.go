package main

import "fmt"

func main() {

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println("Example 1:")
	fmt.Println("Expected outuput: 4")
	fmt.Println("Actual outuput:", search(nums, target))

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 3
	fmt.Println("Example 2:")
	fmt.Println("Expected outuput: -1")
	fmt.Println("Actual outuput:", search(nums, target))
}

func search(nums []int, target int) int {

	n := len(nums)
	low, hi := 0, n-1
	// find a rotation index
	for low < hi {
		mid := low + (hi-low)/2
		if nums[mid] > nums[hi] {
			low = mid + 1
		} else {
			hi = mid
		}
	}

	// find a target by a binary search
	pivot := low
	low, hi = pivot, pivot-1+n
	for low <= hi {
		mid := low + (hi-low)/2
		if nums[mid%n] > target {
			hi = mid - 1
		} else if nums[mid%n] < target {
			low = mid + 1
		} else {
			return mid % n
		}
	}
	return -1
}
