package medium

import "sort"

// Time: O(n*log(n))
// Space: O(n)
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		prev := merged[len(merged)-1]

		if prev[1] >= curr[0] {
			merged[len(merged)-1][1] = max(prev[1], curr[1])
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}
