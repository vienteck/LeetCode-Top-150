package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort intervals based on the start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		previous := result[len(result)-1]

		if current[0] <= previous[1] {
			// Merge overlapping intervals
			result[len(result)-1][1] = max(current[1], previous[1])
		} else {
			// Add non-overlapping interval to the result
			result = append(result, current)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {1, 10}}
	result := merge(intervals)
	fmt.Println(result)
}
