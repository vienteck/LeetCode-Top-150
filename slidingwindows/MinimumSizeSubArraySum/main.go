package main

//Given an array of positive integers nums and
//a positive integer target, return the minimal length of a subarray
//whose sum is greater than or equal to target.
//If there is no such subarray, return 0 instead.

import "math"

func main() {
	minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
}

func minSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt32
	sum := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}

	}
	if minLen == math.MaxInt32 {
		return 0
	}

	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
