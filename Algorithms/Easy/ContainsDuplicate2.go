package easy

// Time: O(n)
// Space: O(n)
func containsNearbyDuplicate(nums []int, k int) bool {
	abs := func(num int) int {
		if num < 0 {
			return -num
		}

		return num
	}

	seen := make(map[int]int, len(nums))
	for i, num := range nums {
		if idx, ok := seen[num]; ok {
			if abs(i-idx) <= k {
				return true
			}
		}

		seen[num] = i
	}

	return false
}