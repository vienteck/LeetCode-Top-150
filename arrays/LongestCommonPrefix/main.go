package main

func main() {
	longestCommonPrefix([]string{"ab", "a"})
}

func longestCommonPrefix(strs []string) string {
	if len(strs) < 2 {
		return strs[0]
	}
	word := strs[0]
	strs = strs[1:]
	lcp := ""
	for j := 0; j < len(word); j++ {
		add := true
		for i := 0; i < len(strs); i++ {
			if j > len(strs[i])-1 || word[j] != strs[i][j] {
				add = false
				break
			}
		}
		if !add {
			break
		} else {
			lcp += string(word[j])
		}
	}
	return lcp
}
