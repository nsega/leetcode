package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {

	countMap := make(map[int]int)
	var sum int
	var count int

	for _, num := range nums {

		sum += num
		if sum == k {
			count += 1
		}

		if v, ok := countMap[sum-k]; ok {
			count += v
		}

		if v, ok := countMap[sum]; !ok {
			countMap[sum] = 1
		} else {
			countMap[sum] = v + 1
		}
	}
	return count
}

func main() {
	nums := []int{1, 1, 1}
	k := 2
	fmt.Println(subarraySum(nums, k))
}
