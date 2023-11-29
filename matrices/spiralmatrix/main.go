package main

func main() {
	matrix := [][]int{{1, 11}, {2, 12}, {3, 13}, {4, 14}, {5, 15}, {6, 16}, {7, 17}, {8, 18}, {9, 19}, {10, 20}}

	spiralOrder(matrix)
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	results := []int{}

	total := len(matrix) * len(matrix[0])
	processed := 0

	for processed < total || len(matrix) > 0 {
		if len(matrix) > 0 {
			for i := 0; i < len(matrix[0]); i++ {
				results = append(results, matrix[0][i])
				processed++
			}
			matrix = matrix[1:]
		}

		if len(matrix) > 0 {
			for i := 0; i < len(matrix); i++ {
				n := len(matrix[i]) - 1
				results = append(results, matrix[i][n])
				if n == 0 {
					matrix = matrix[1:]
					i--
				} else {
					matrix[i] = matrix[i][:n]
				}
				processed++
			}

		}

		if len(matrix) > 0 {
			n := len(matrix) - 1
			j := len(matrix[n]) - 1
			for i := j; i >= 0; i-- {
				results = append(results, matrix[n][i])
				processed++
			}
			matrix = matrix[:n]
		}

		if len(matrix) > 0 {
			for i := len(matrix) - 1; i >= 0; i-- {
				n := len(matrix[i]) - 1
				results = append(results, matrix[i][0])
				if n == 0 {
					matrix = matrix[:i]

				} else {
					matrix[i] = matrix[i][1:]
				}
				processed++
			}

		}
	}
	return results
}
