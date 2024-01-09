package main

func main() {

}

func singleNumber(nums []int) int {
	tracker := map[int]int{}

	for _, val := range nums {
		if _, exists := tracker[val]; exists {
			tracker[val]++
		} else {
			tracker[val] = 1
		}
	}
	for key, value := range tracker {
		if value == 1 {
			return key
		}
	}
	return 0
}
