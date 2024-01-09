package main

func main() {
	getRow(3)
}

func getRow(rowIndex int) []int {
	answer := [][]int{[]int{1}, []int{1, 1}}
	if rowIndex == 0 {
		return answer[0]
	}
	if rowIndex == 1 {
		return answer[1]
	}

	for i := 2; i <= rowIndex; i++ {
		answer = append(answer, make([]int, i+1))
		for j := 0; j < len(answer[i]); j++ {
			if j == 0 || j == len(answer[i])-1 {
				answer[i][j] = 1
			} else {
				answer[i][j] = answer[i-1][j-1] + answer[i-1][j]
			}
		}
	}
	return answer[len(answer)-1]
}
