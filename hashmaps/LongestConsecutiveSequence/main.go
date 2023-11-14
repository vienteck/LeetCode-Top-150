package main

func main() {
	longestConsecutive([]int{100, 3, 200, 1, 3, 2})
}

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)

	// Populate the set
	for _, num := range nums {
		numSet[num] = true
	}

	maxLen := 0

	// Iterate through the array
	for _, num := range nums {
		// Check if it's the start of a sequence
		if _, exists := numSet[num-1]; !exists {
			currentNum := num
			currentLen := 1

			// Check for consecutive elements in the set
			for numSet[currentNum+1] {
				currentNum++
				currentLen++
			}

			// Update the maximum length
			if currentLen > maxLen {
				maxLen = currentLen
			}
		}
	}

	return maxLen
}
