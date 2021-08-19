package main

import "fmt"

func numIslands(grid [][]byte) int {
	maxY := len(grid)
	if maxY == 0 {
		return 0
	}
	maxX := len(grid[0])

	visited := make([][]int, maxY)
	for i := 0; i < maxY; i++ {
		visited[i] = make([]int, maxX)
	}

	res := 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if grid[y][x] == '1' && visited[y][x] == 0 {
				res++
				fill(grid, visited, y, x)
			}
		}
	}
	return res
}

func fill(grid [][]byte, visited [][]int, y, x int) {
	visited[y][x] = 1
	maxY, maxX := len(grid), len(grid[y])
	if y > 0 && grid[y-1][x] == '1' && visited[y-1][x] == 0 {
		fill(grid, visited, y-1, x)
	}
	if x < maxX-1 && grid[y][x+1] == '1' && visited[y][x+1] == 0 {
		fill(grid, visited, y, x+1)
	}
	if y < maxY-1 && grid[y+1][x] == '1' && visited[y+1][x] == 0 {
		fill(grid, visited, y+1, x)
	}
	if x > 0 && grid[y][x-1] == '1' && visited[y][x-1] == 0 {
		fill(grid, visited, y, x-1)
	}
}

func main() {
	type args struct {
		grid [][]byte
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
				grid: [][]byte{
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			exp: 1,
		},
		{
			name: "Example 2:",
			arg: args{
				grid: [][]byte{
					{'1', '1', '0', '0', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '1', '0', '0'},
					{'0', '0', '0', '1', '1'},
				},
			},
			exp: 3,
		},
	}

	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println(" Actual: ", numIslands(tt.arg.grid))
		fmt.Println(" Exp: ", tt.exp)
	}
}
