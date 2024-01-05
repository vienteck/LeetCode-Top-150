func minOperations(nums []int) int {
	occurMap := make(map[int]int)

	for _, val := range nums {
		occurMap[val]++
	}

	operations := 0
	for _, val := range occurMap {
		if val == 1 {
			return -1
		}
		operations += int(math.Ceil(float64(val) / 3.0))
	}

	return operations
}