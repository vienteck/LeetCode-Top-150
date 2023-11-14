package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		// Convert the string to a sorted string, which will be the key for anagrams
		sortedStr := sortString(str)

		// Append the original string to the corresponding anagram group
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	// Convert the map values to a 2D slice
	var result [][]string
	for _, v := range anagrams {
		result = append(result, v)
	}

	return result
}

func sortString(s string) string {
	// Convert the string to a slice of characters, sort it, and join back into a string
	sChars := strings.Split(s, "")
	sort.Strings(sChars)
	return strings.Join(sChars, "")
}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output := groupAnagrams(input)
	fmt.Println(output)
}
