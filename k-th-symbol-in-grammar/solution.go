package main

import (
	"fmt"
	"math/bits"
)

func main() {
	N := 1
	K := 1
	fmt.Printf("Example 1: \n")
	fmt.Printf("  Input: %d, %d\n", N, K)
	fmt.Printf("  Expected output: %d\n", 0)
	fmt.Printf("  Actual output: %d\n", kthGrammar(N, K))

	N = 2
	K = 1
	fmt.Printf("Example 2: \n")
	fmt.Printf("  Input: %d, %d\n", N, K)
	fmt.Printf("  Expected output: %d\n", 0)
	fmt.Printf("  Actual output: %d\n", kthGrammar(N, K))

	N = 2
	K = 2
	fmt.Printf("Example 3: \n")
	fmt.Printf("  Input: %d, %d\n", N, K)
	fmt.Printf("  Expected output: %d\n", 1)
	fmt.Printf("  Actual output: %d\n", kthGrammar(N, K))

	N = 4
	K = 5
	fmt.Printf("Example 4: \n")
	fmt.Printf("  Input: %d, %d\n", N, K)
	fmt.Printf("  Expected output: %d\n", 1)
	fmt.Printf("  Actual output: %d\n", kthGrammar(N, K))
}

func kthGrammar(N int, K int) int {
	return bits.OnesCount(uint(K-1)) & 1
}
