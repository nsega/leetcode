package main

import "fmt"

func main() {

	nums := []int{1, 2, 3}
	exp := []int{1, 3, 2}
	fmt.Println("Example 1:")
	fmt.Printf("  Input: nums = %v\n", nums)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", nextPermutation(nums))

	nums = []int{3, 2, 1}
	exp = []int{1, 2, 3}
	fmt.Println("Example 2:")
	fmt.Printf("  Input: nums = %v\n", nums)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", nextPermutation(nums))

	nums = []int{1, 1, 5}
	exp = []int{1, 5, 1}
	fmt.Println("Example 3:")
	fmt.Printf("  Input: nums = %v\n", nums)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", nextPermutation(nums))
}

func nextPermutation(nums []int) []int {
	i := -1
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i == -1 {
		reverseSort(nums)
		return nums
	}

	j := 0
	for j = len(nums) - 1; j > i; j-- {
		if nums[j] > nums[i] {
			break
		}
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverseSort(nums[i+1:])
	return nums
}

func reverseSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		j := len(nums) - 1 - i
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
}
