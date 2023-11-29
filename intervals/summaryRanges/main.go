package main

import (
	"fmt"
	"strconv"
)

func main() {
	summaryRanges([]int{0, 1, 2, 4, 5, 7})
}

func summaryRanges(nums []int) []string {
	results := []string{}
	start := nums[0]
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-prev > 1 {
			p := strconv.Itoa(prev)
			s := strconv.Itoa(start)
			if start == prev {
				results = append(results, p)
			} else {
				results = append(results, fmt.Sprintf("%v->%v", s, p))
			}
			start = nums[i]
		} else {
			prev = nums[i]
		}
	}
	return results
}
