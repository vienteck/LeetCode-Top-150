package main

import "strings"

func main() {
	reverseWords("the sky is blue")
}

func reverseWords(s string) string {
	container := []string{}
	word := ""
	for _, r := range s {
		if string(r) == " " {
			if word != "" {
				container = append([]string{word}, container...)
				word = ""
			}
		} else {
			word += string(r)
		}
	}

	if word != "" {
		container = append([]string{word}, container...)
	}

	return strings.Join(container, " ")
}
