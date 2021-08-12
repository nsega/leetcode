package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	nums1map := make(map[int]bool)
	for _, n := range nums1 {
		nums1map[n] = true
	}

	nums2map := make(map[int]bool)
	for _, n := range nums2 {
		nums2map[n] = true
	}

	var res []int
	for k := range nums2map {
		if _, ok := nums1map[k]; ok {
			res = append(res, k)
		}
	}
	return res
}

func main() {
	type args struct {
		nums1 []int
		nums2 []int
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
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			exp: []int{2},
		},
		{
			name: "Example 2:",
			arg: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			exp: []int{4, 9},
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual:", intersection(tt.arg.nums1, tt.arg.nums2))
		fmt.Println(" Exp:", tt.exp)
	}
}
