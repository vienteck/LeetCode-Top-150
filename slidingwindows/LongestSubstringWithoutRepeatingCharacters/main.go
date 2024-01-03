package main

func main() {
	slidingWindowVersion("abcabcbb")
}

func slidingWindowVersion(s string) int {
	charIndex := make(map[byte]int)
	start, maxLength := 0, 0

	for end := 0; end < len(s); end++ {
		if index, exists := charIndex[s[end]]; exists && index >= start {
			start = index + 1
		}
		charIndex[s[end]] = end
		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}
	answer := 1

	for i := 0; i < len(s); i++ {
		tracker := map[string]struct{}{}
		tracker[string(s[i])] = struct{}{}
		maxLength := 1
		for j := i + 1; j < len(s); j++ {
			if _, exists := tracker[string(s[j])]; !exists {
				tracker[string(s[j])] = struct{}{}
				maxLength++
			} else {
				break
			}
			answer = getmax(answer, maxLength)
		}
	}

	return answer
}

func getmax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
