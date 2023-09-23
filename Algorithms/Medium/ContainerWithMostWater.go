package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(1)
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	max_area := 0

	for l < r {
		area := (r - l) * helpers.Min(height[l], height[r])
		max_area = helpers.Max(max_area, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return max_area
}
