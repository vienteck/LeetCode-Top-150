package main

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
}

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		reverse(matrix[i])
	}
}

func reverse(row []int) {
	i, j := 0, len(row)-1
	for i < j {
		row[i], row[j] = row[j], row[i]
		i++
		j--
	}
}
