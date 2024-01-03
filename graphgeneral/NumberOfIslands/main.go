package main

import "fmt"

func main() {
	// Example usage
	grid := [][]byte{
		[]byte("11110"),
		[]byte("11010"),
		[]byte("11000"),
		[]byte("00000"),
	}

	result := numIslands(grid)
	fmt.Println(result)
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == '1' {
			grid[i][j] = '0' // Mark the current cell as visited
			// Explore the neighbors in all four directions
			dfs(i+1, j)
			dfs(i-1, j)
			dfs(i, j+1)
			dfs(i, j-1)
		}
	}

	islandCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				islandCount++
				dfs(i, j)
			}
		}
	}

	return islandCount
}
