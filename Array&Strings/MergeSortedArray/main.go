package main

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := []int{}
	l, r := 0, 0
	for l < m && r < n {
		if nums1[l] < nums2[r] {
			sorted = append(sorted, nums1[l])
			l++
		} else {
			sorted = append(sorted, nums2[r])
			r++
		}
	}

	if l == m {
		sorted = append(sorted, nums2[r:]...)
	} else {
		sorted = append(sorted, nums1[l:]...)
	}

	copy(nums1, sorted[:m+n])
}
