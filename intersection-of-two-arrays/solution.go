package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {

	nums1map := make(map[int]bool)
	for _, num := range nums1 {
		nums1map[num] = true
	}

	nums2map := make(map[int]bool)
	for _, num := range nums2 {
		nums2map[num] = true
	}

	var result []int
	for k := range nums2map {
		if _, ok := nums1map[k]; ok {
			result = append(result, k)
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}
