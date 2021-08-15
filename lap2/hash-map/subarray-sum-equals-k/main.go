package main

import "fmt"

// subarraySum is to return the total number of continuous
// subarrays whose sum equals to k.
func subarraySum(nums []int, k int) int {
	countMap := make(map[int]int)
	sum, count := 0, 0
	for _, num := range nums {
		sum = sum + num
		if sum == k {
			count++
		}
		if v, ok := countMap[sum-k]; ok {
			count = count + v
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
	type args struct {
		nums []int
		k    int
	}
	type test struct {
		name string
		arg  args
		exp  int
	}
	tests := []test{
		{
			name: "Example 1:",
			arg: args{
				nums: []int{1, 1, 1},
				k:    2,
			},
			exp: 2,
		},
		{
			name: "Example 2:",
			arg: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			exp: 2,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual:", subarraySum(tt.arg.nums, tt.arg.k))
		fmt.Println(" Exp:", tt.exp)
	}
}
