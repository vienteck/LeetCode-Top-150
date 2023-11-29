func isSubsequence(s string, t string) bool {
	for _, val := range t {
		if len(s) > 0 {
			if val == rune(s[0]) {
				s = s[1:]
			}
		}
	}
	if len(s) > 0 {
		return false
	} else {
		return true
	}
}