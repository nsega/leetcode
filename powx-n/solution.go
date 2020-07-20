package main

import "fmt"

func main() {
	x := 2.00000
	n := 10
	fmt.Println("Example 1:")
	fmt.Printf("  Input: %f, %d\n", x, n)
	fmt.Println("  Expected output: ", 1024.00000)
	fmt.Println("  Actual output: ", myPow(x, n))

	x = 2.10000
	n = 3
	fmt.Println("Example 2:")
	fmt.Printf("  Input: %f, %d\n", x, n)
	fmt.Println("  Expected output: ", 9.2610000)
	fmt.Println("  Actual output: ", myPow(x, n))

	x = 2.00000
	n = -2
	fmt.Println("Example 3:")
	fmt.Printf("  Input: %f, %d\n", x, n)
	fmt.Println("  Expected output: ", 0.25000)
	fmt.Println("  Actual output: ", myPow(x, n))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	}
	return x * myPow(x*x, n/2)
}
