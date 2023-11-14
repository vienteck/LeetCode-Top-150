package main

import "strings"

func main() {
	wordPattern("abba", "dog dog dog dog")
}

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	pToS := map[byte]string{}
	sToP := map[string]byte{}

	for i := 0; i < len(pattern); i++ {
		pvalue, pexists := pToS[pattern[i]]
		svalue, sexists := sToP[words[i]]

		if pexists {
			if pvalue != words[i] {
				return false
			}
		} else {
			pToS[pattern[i]] = words[i]
		}

		if sexists {
			if svalue != pattern[i] {
				return false
			}
		} else {
			sToP[words[i]] = pattern[i]
		}
	}

	return true
}
