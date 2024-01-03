package main

func main() {
	findMatrix([]int{1, 3, 4, 1, 2, 3, 1})
}

func findMatrix(nums []int) [][]int {

	if len(nums) < 2 {
		answer := make([][]int, 1)
		answer[0] = append(answer[0], nums...)
		return answer
	}
	tracker := make([]map[int]struct{}, 1)
	tracker[0] = map[int]struct{}{}
	for len(nums) > 0 {
		current := nums[0]
		added := false
		//check if number exists in lists
		for key, val := range tracker {
			if _, exists := val[current]; !exists {
				// we need to add number to this list
				tracker[key][current] = struct{}{}
				added = true
				break
			}
		}

		if !added {
			tracker = append(tracker, map[int]struct{}{})
			n := len(tracker)
			tracker[n-1] = map[int]struct{}{}
			tracker[n-1][current] = struct{}{}
		}
		nums = nums[1:]
	}
	answer := make([][]int, len(tracker))

	for index, mapvalue := range tracker {
		for key := range mapvalue {
			answer[index] = append(answer[index], key)
		}
	}

	return answer
}
