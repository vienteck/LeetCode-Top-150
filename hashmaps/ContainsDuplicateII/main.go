package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	numIndex := make(map[int]int)

	for i, num := range nums {
		if index, ok := numIndex[num]; ok && i-index <= k {
			return true
		}
		numIndex[num] = i
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3
	result := containsNearbyDuplicate(nums, k)
	fmt.Println(result) // Output: true
}
