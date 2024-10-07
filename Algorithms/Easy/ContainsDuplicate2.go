package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time: O(n)
// Space: O(n)
func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int, len(nums))
	for i, num := range nums {
		if idx, ok := seen[num]; ok {
			if helpers.Abs(i-idx) <= k {
				return true
			}
		}

		seen[num] = i
	}

	return false
}
