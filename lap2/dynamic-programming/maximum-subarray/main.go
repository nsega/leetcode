package main

import "fmt"

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	type args struct {
		nums []int
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
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			exp: 6,
		},
		{
			name: "Example 2:",
			arg: args{
				nums: []int{1},
			},
			exp: 1,
		},
		{
			name: "Example 3:",
			arg: args{
				nums: []int{5, 4, -1, 7, 8},
			},
			exp: 23,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual:", maxSubArray(tt.arg.nums))
		fmt.Println(" Exp:", tt.exp)
	}
}
