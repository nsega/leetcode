package main

import "fmt"

func main() {
	nums := []int{0,1,0,3,12}
	exp := []int{1,3,12,0,0}
	fmt.Println("Example :")
	fmt.Printf("  Input: nums = %v\n", nums)
	fmt.Println("  Expected output: ", exp)
	fmt.Println("  Actual output: ", moveZeroes(nums))
}

func moveZeroes(nums []int) []int{
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return nums
}