package main

import "fmt"

func canJump(nums []int) bool {
	maxReach := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > maxReach {
			return false // If you can't reach the current position, return false
		}

		maxReach = max(maxReach, i+nums[i])

		if maxReach >= n-1 {
			return true // If you can reach the last index, return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example usage:
	nums := []int{2, 3, 1, 1, 4}
	result := canJump(nums)
	fmt.Println(result)
}
