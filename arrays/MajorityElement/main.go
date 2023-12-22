package main

func main() {

}

func majorityElement(nums []int) int {
	target := len(nums) / 2
	tracker := map[int]int{}
	for _, val := range nums {
		if _, exists := tracker[val]; exists {
			tracker[val]++
			if tracker[val] >= target {
				return val
			}
		} else {
			tracker[val] = 1
		}
	}

	return -1
}
