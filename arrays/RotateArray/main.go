package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	// Reverse the entire array
	reverse(nums)

	// Reverse the first k elements
	reverse(nums[:k])

	// Reverse the remaining elements
	reverse(nums[k:])
}

func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {
	// Example usage:
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 5
	rotate(nums, k)
	fmt.Println(nums)
}
