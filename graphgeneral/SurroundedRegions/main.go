package main

import "fmt"

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	// Helper function to perform depth-first search (DFS)
	var dfs func(row, col int)
	dfs = func(row, col int) {
		// Check if the current cell is out of bounds or not 'O'
		if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != 'O' {
			return
		}

		// Mark the current cell as 'T' (temporary mark)
		board[row][col] = 'T'

		// Directions for DFS: right, left, down, up
		directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, dir := range directions {
			nr, nc := row+dir[0], col+dir[1]
			dfs(nr, nc)
		}
	}

	// Mark 'O' connected to the borders as 'T'
	for i := 0; i < rows; i++ {
		dfs(i, 0)
		dfs(i, cols-1)
	}

	for j := 0; j < cols; j++ {
		dfs(0, j)
		dfs(rows-1, j)
	}

	// Capture surrounded 'O' and revert 'T' back to 'O'
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}

func main() {
	// Example usage
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	fmt.Println("Original Board:")
	printBoard(board)

	solve(board)

	fmt.Println("\nBoard after solving:")
	printBoard(board)
}

func printBoard(board [][]byte) {
	for _, row := range board {
		fmt.Println(string(row))
	}
}
