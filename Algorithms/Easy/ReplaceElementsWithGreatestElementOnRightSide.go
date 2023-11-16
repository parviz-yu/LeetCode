package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(1)
func replaceElements(arr []int) []int {
	maxVal := -1
	for i := len(arr) - 1; i > -1; i-- {
		temp := arr[i]
		arr[i] = maxVal
		maxVal = helpers.Max(temp, maxVal)
	}

	return arr
}
