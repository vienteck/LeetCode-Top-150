package main

import "fmt"

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '.', '.', '5', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'9', '3', '.', '.', '2', '.', '4', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '3', '4', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '5', '2', '.', '.'},
	}
	isValidSudoku(board)
}

func isValidSudoku(board [][]byte) bool {
	//Check rows
	for i := 0; i < 9; i++ {
		tracker := map[byte]struct{}{}
		for j := 0; j < 9; j++ {
			if board[i][j] == 46 {
				continue
			}
			if _, exists := tracker[board[i][j]]; exists {
				return false
			} else {
				tracker[board[i][j]] = struct{}{}
			}
		}
	}
	//check columns
	for i := 0; i < 9; i++ {
		tracker := map[byte]struct{}{}
		for j := 0; j < 9; j++ {
			if board[j][i] == 46 {
				continue
			}
			if _, exists := tracker[board[j][i]]; exists {
				return false
			} else {
				tracker[board[j][i]] = struct{}{}
			}
		}
	}
	//check each 3x3 grid
	for i := 0; i < 3; i++ {
		m := map[int]int{
			0: 0,
			1: 3,
			2: 6,
		}
		tracker := map[byte]struct{}{}
		for h := 0; h < 9; h++ {
			for j := m[i]; j < m[i]+3; j++ {
				fmt.Println(string(board[h][j]))
				if h%3 == 0 && h != 0 && (j == 0 || j == 3 || j == 6) {
					tracker = map[byte]struct{}{}
				}
				if board[h][j] == 46 {
					continue
				}
				if _, exists := tracker[board[h][j]]; exists {
					return false
				} else {
					tracker[board[h][j]] = struct{}{}
				}
			}
		}

	}
	return true
}
