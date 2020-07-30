package main

import (
	"fmt"
	"math"
)

func main() {

	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println("Example:")
	fmt.Printf("  Input: s = %d nums = %v\n", s, nums)
	fmt.Println("  Expected output: ", 2)
	fmt.Println("  Actual output: ", minSubArrayLen(s, nums))

}

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	j := 0
	sum := 0
	ret := math.MaxInt32

	for j < len(nums) {
		sum += nums[j]
		j++
		for sum >= s {
			ret = min(ret, j-i)
			sum -= nums[i]
			i++
		}
	}
	if ret == math.MaxInt32 {
		return 0
	}
	return ret
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
