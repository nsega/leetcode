package main

import (
	"fmt"
	"sort"
)

func searchInsert(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
}

func main() {
	type args struct {
		nums   []int
		target int
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
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			exp: 2,
		},
		{
			name: "Example 2:",
			arg: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			exp: 1,
		},
		{
			name: "Example 3:",
			arg: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			exp: 4,
		},
		{
			name: "Example 4:",
			arg: args{
				nums:   []int{1, 3, 5, 6},
				target: 0,
			},
			exp: 0,
		},
		{
			name: "Example 5:",
			arg: args{
				nums:   []int{1},
				target: 0,
			},
			exp: 0,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual:", searchInsert(tt.arg.nums, tt.arg.target))
		fmt.Println(" Exp:", tt.exp)
	}
}
