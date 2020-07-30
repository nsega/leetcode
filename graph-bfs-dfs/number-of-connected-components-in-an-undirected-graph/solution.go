package main

import (
	"fmt"
)

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
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
	fmt.Println(countComponents(n, edges))

	n = 5
	edges = [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
	fmt.Println(countComponents(n, edges))
}
