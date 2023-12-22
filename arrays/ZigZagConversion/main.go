package main

import "strings"

func main() {
	convert("PAYPALISHIRING", 3)
}

func convert(s string, numRows int) string {

	var container [][]string
	current := 0
	adder := -1
	for i := 0; i < numRows; i++ {
		container = append(container, []string{})
	}

	for len(s) > 0 {
		container[current] = append(container[current], string(s[0]))
		if current == numRows-1 || current == 0 {
			adder *= -1
		}
		current += adder
		s = s[1:]
	}
	result := ""
	for i := 0; i < numRows; i++ {
		result += strings.Join(container[i], "")
	}
	return ""
}
