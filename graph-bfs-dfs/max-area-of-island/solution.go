package main

import (
	"fmt"
)

func maxAreaOfIsland(grid [][]int) int {

	maxY := len(grid)
	if maxY == 0 {
		return 0
	}
	maxX := len(grid[0])

	maxArea := 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if grid[y][x] > 0 {
				if sum := sumIslandArea(grid, y, x); sum > maxArea {
					maxArea = sum
				}
			}
		}
	}
	return maxArea

}

func sumIslandArea(grid [][]int, y, x int) int {

	sum := 1
	grid[y][x] = 0

	maxY, maxX := len(grid), len(grid[y])

	if y > 0 && grid[y-1][x] > 0 {
		sum += sumIslandArea(grid, y-1, x)
	}
	if x < maxX-1 && grid[y][x+1] > 0 {
		sum += sumIslandArea(grid, y, x+1)
	}
	if y < maxY-1 && grid[y+1][x] > 0 {
		sum += sumIslandArea(grid, y+1, x)
	}
	if x > 0 && grid[y][x-1] > 0 {
		sum += sumIslandArea(grid, y, x-1)
	}
	return sum
}

func main() {

	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))

}
