package main

import "fmt"

// numWays is return the number of ways you can paint the fence.
//
// You are painting a fence of n posts with k different colors.
// You must paint the posts following these rules:
// - Every post must be painted exactly one color.
// - There cannot be three or more consecutive posts with the same color.
func numWays(n int, k int) int {
	if n == 0 || (n > 2 && k == 1) {
		return 0
	}
	if n == 1 {
		return k
	}
	if n == 2 {
		return k + k*(k-1)
	}
	dp := make([]int, n)
	dp[0] = k
	dp[1] = k*(k-1) + k
	for i := 2; i < n; i++ {
		dp[i] = (k - 1) * (dp[i-1] + dp[i-2])
	}
	return dp[n-1]
}

func main() {
	type args struct {
		n int
		k int
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
				n: 3,
				k: 2,
			},
			exp: 6,
		},
		{
			name: "Example 2:",
			arg: args{
				n: 1,
				k: 1,
			},
			exp: 1,
		},
		{
			name: "Example 3:",
			arg: args{
				n: 7,
				k: 2,
			},
			exp: 42,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual:", numWays(tt.arg.n, tt.arg.k))
		fmt.Println(" Exp:", tt.exp)
	}
}
