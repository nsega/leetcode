package main

import "fmt"

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

func main() {
	type args struct {
		x float64
		n int
	}
	type test struct {
		name string
		arg  args
		exp  float64
	}
	tests := []test{
		{
			name: "Example 1:",
			arg: args{
				x: 2.00000,
				n: 10,
			},
			exp: 1024.00000,
		},
		{
			name: "Example 2:",
			arg: args{
				x: 2.10000,
				n: 3,
			},
			exp: 9.26100,
		},
		{
			name: "Example 3:",
			arg: args{
				x: 2.00000,
				n: -2,
			},
			exp: 0.25000,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual:", myPow(tt.arg.x, tt.arg.n))
		fmt.Println(" Exp:", tt.exp)
	}
}
