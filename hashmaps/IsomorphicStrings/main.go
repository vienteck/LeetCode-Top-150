package main

func main() {
	isIsomorphic("paper", "title")
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		charS, existsS := sToT[s[i]]
		charT, existsT := tToS[t[i]]

		// Check if the mapping is consistent
		if existsS {
			if charS != t[i] {
				return false
			}
		} else {
			sToT[s[i]] = t[i]
		}

		if existsT {
			if charT != s[i] {
				return false
			}
		} else {
			tToS[t[i]] = s[i]
		}
	}

	return true
}
