package main

func main() {
	removeDuplicates([]int{1, 1, 2})
}

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}

func removeDuplicatesFirstSolution(nums []int) int {
	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			k++
		} else {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}

	}
	return k
}
