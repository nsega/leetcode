package main

import "fmt"

func twoSum(nums []int, target int) []int {
	matched := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == (nums[i] + nums[j]) {
				matched = append(matched, i)
				matched = append(matched, j)
				break
			}
		}
	}
	return matched
}

func main() {
	type args struct {
		nums   []int
		target int
	}
	type test struct {
		name string
		arg  args
		exp  []int
	}
	tests := []test{
		{
			name: "Example 1:",
			arg: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			exp: []int{0, 1},
		},
		{
			name: "Example 2:",
			arg: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			exp: []int{1, 2},
		},
		{
			name: "Example 3:",
			arg: args{
				nums:   []int{3, 3},
				target: 6,
			},
			exp: []int{0, 1},
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual: ", twoSum(tt.arg.nums, tt.arg.target))
		fmt.Println(" Exp: ", tt.exp)
	}
}
