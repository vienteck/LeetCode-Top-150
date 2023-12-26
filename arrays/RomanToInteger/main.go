package main

func main() {

}

func romanToInt(s string) int {
	answer := 0
	prev := 0
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	for i := len(s) - 1; i >= 0; i-- {
		if m[string(s[i])] < prev {
			answer -= m[string(s[i])]
		} else {
			answer += m[string(s[i])]
		}
		prev = m[string(s[i])]
	}

	return answer
}
