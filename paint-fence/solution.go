package main

import "fmt"

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
	n := 3
	k := 2
	fmt.Println("output:", numWays(n, k))
}
