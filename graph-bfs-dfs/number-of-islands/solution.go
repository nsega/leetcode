package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	mlen := len(grid)
	if mlen == 0 {
		return 0
	}
	nlen := len(grid[0])

	visited := make([][]int, mlen)
	for i := 0; i < mlen; i++ {
		visited[i] = make([]int, nlen)
	}

	res := 0
	for i := 0; i < mlen; i++ {
		for j := 0; j < nlen; j++ {
			if grid[i][j] == '1' && visited[i][j] == 0 {
				res++
				fill(grid, visited, i, j, mlen, nlen)
			}
		}
	}
	return res
}

func fill(grid [][]byte, visited [][]int, i, j int, mlen, nlen int) {

	visited[i][j] = 1
	if i-1 >= 0 && grid[i-1][j] == '1' && visited[i-1][j] == 0 {
		fill(grid, visited, i-1, j, mlen, nlen)
	}

	if i+1 < mlen && grid[i+1][j] == '1' && visited[i+1][j] == 0 {
		fill(grid, visited, i+1, j, mlen, nlen)
	}

	if j-1 >= 0 && grid[i][j-1] == '1' && visited[i][j-1] == 0 {
		fill(grid, visited, i, j-1, mlen, nlen)
	}

	if j+1 < nlen && grid[i][j+1] == '1' && visited[i][j+1] == 0 {
		fill(grid, visited, i, j+1, mlen, nlen)
	}
}

func main() {
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(grid))
}
