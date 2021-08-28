package main

import "fmt"

func countComponents(n int, edges [][]int) int {
	roots := make([]int, n)
	for i := 0; i < n; i++ {
		roots[i] = i
	}
	for _, edge := range edges {
		root1 := find(roots, edge[0])
		root2 := find(roots, edge[1])
		if root1 != root2 {
			roots[root1] = root2
			n--
		}
	}
	return n
}

func find(roots []int, ID int) int {
	for roots[ID] != ID {
		roots[ID] = roots[roots[ID]]
		ID = roots[ID]
	}
	return ID
}

func main() {
	type args struct {
		n     int
		edges [][]int
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
				n:     5,
				edges: [][]int{{0, 1}, {1, 2}, {3, 4}},
			},
			exp: 2,
		},
		{
			name: "Example 2:",
			arg: args{
				n:     5,
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			},
			exp: 1,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual: ", countComponents(tt.arg.n, tt.arg.edges))
		fmt.Println(" Exp: ", tt.exp)
	}
}
