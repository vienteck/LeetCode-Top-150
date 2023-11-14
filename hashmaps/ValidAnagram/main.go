package main

func main() {
	isAnagram("anagram", "nagaram")
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	tracker := map[rune]int{}
	//build out the tracker with all the runes and their counts from strings
	for _, x := range s {
		if _, ok := tracker[x]; ok {
			tracker[x]++
		} else {
			tracker[x] = 1
		}
	}
	//after iterating through the tracker, it should have a len of 0
	for _, x := range t {
		if _, ok := tracker[x]; ok {
			if tracker[x] > 0 {
				tracker[x]--
				if tracker[x] == 0 {
					delete(tracker, x)
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return len(tracker) == 0
}
