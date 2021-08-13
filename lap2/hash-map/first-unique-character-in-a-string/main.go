package main

import "fmt"

// firstUniqChar is to find the first non-repeating character in it and return its index.
// If it does not exist, return -1.
func firstUniqChar(s string) int {
	flags := make([]int, 26)
	for i := range flags {
		flags[i] = -1
	}
	slen := len(s)
	for i, ch := range s {
		idx := ch - 'a'
		if flags[idx] == -1 {
			flags[idx] = i
		} else {
			flags[idx] = slen
		}
	}
	min := slen
	for i := range flags {
		if flags[i] >= 0 && flags[i] < len(s) && flags[i] < min {
			min = flags[i]
		}
	}
	if min == slen {
		return -1
	}
	return min
}

func main() {
	type args struct {
		s string
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
				s: "leetcode",
			},
			exp: 0,
		},
		{
			name: "Example 2:",
			arg: args{
				s: "loveleetcode",
			},
			exp: 2,
		},
		{
			name: "Example 3:",
			arg: args{
				s: "aabb",
			},
			exp: -1,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual:", firstUniqChar(tt.arg.s))
		fmt.Println(" Exp:", tt.exp)
	}
}
