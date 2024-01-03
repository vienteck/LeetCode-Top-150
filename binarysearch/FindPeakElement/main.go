package main

func main() {

}

func findPeakElement(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}

	if nums[0] > nums[len(nums)-1] {
		return 0
	} else {
		return len(nums) - 1
	}
}
