package medium

// Time: O(n)
// Space: O(n)
func pivotArray(nums []int, pivot int) []int {
	less := make([]int, 0)
	equal := make([]int, 0)
	greater := make([]int, 0)

	for _, num := range nums {
		if num < pivot {
			less = append(less, num)
		} else if num == pivot {
			equal = append(equal, num)
		} else {
			greater = append(greater, num)
		}
	}

	return append(less, append(equal, greater...)...)
}
