package main

import (
	"fmt"
)

func convertToTitle(columnNumber int) string {
	result := ""

	for columnNumber > 0 {
		columnNumber-- // Adjust to 0-based index
		mod := columnNumber % 26
		result = string('A'+mod) + result
		columnNumber /= 26
	}

	return result
}

func main() {
	// Example usage:
	columnNumber := 168
	result := convertToTitle(columnNumber)
	fmt.Println(result)
}
