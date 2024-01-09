package main

func main() {
	generate(5)
}

func generate(numRows int) [][]int {
	answer := [][]int{[]int{1}, []int{1, 1}}
	if numRows == 1 {
		return [][]int{[]int{1}}
	}
	if numRows == 2 {
		return answer
	}

	for i := 2; i < numRows; i++ {
		answer = append(answer, make([]int, i+1))
		for j := 0; j < len(answer[i]); j++ {
			if j == 0 || j == len(answer[i])-1 {
				answer[i][j] = 1
			} else {
				answer[i][j] = answer[i-1][j-1] + answer[i-1][j]
			}
		}
		
	}
	return answer
}
