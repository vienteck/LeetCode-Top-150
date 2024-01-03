package main

import (
	"fmt"
)

var digitMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	var combinations []string

	var backtrack func(index int, path string)
	backtrack = func(index int, path string) {
		if index == len(digits) {
			combinations = append(combinations, path)
			return
		}

		currentDigit := string(digits[index])
		for _, letter := range digitMap[currentDigit] {
			backtrack(index+1, path+string(letter))
		}
	}

	backtrack(0, "")

	return combinations
}

func main() {
	// Example usage
	digits := "233"
	result := letterCombinations(digits)
	fmt.Println(result)
}
