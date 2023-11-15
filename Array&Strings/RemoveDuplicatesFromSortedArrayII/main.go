package main

func main() {
	removeDuplicates([]int{1, 1, 1, 2, 2, 3})
}

func removeDuplicates(nums []int) int {
	tracker := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := tracker[nums[i]]; ok {
			if tracker[nums[i]] > 1 {
				nums = append(nums[:i], nums[i+1:]...)
				i--
			} else {
				tracker[nums[i]]++
			}
		} else {
			tracker[nums[i]] = 1
		}
	}
	return len(nums)
}
