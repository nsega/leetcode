package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("findMin method")
	nums := []int{3,4,5,1,2}
	fmt.Println("Example1:")
	fmt.Println("Expected Output: ", 1)
	fmt.Println("Actual Output: ", findMin(nums))

	nums = []int{4,5,6,7,0,1,2}	
	fmt.Println("Example2:")
	fmt.Println("Expected Output: ", 0)
	fmt.Println("Actual Output: ", findMin(nums))

	fmt.Println("================")
	fmt.Println("findMin2 method")
	nums = []int{3,4,5,1,2}
	fmt.Println("Example1:")
	fmt.Println("Expected Output: ", 1)
	fmt.Println("Actual Output: ", findMin2(nums))

	nums = []int{4,5,6,7,0,1,2}
	fmt.Println("Example2:")
	fmt.Println("Expected Output: ", 0)
	fmt.Println("Actual Output: ", findMin2(nums))

}

func findMin(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

func findMin2(nums []int) int {

	left, right := 0, len(nums) -1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]	
}
