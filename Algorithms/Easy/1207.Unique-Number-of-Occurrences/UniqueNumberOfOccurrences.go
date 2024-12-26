package easy

// Time: O(n)
// Space: O(n)
func uniqueOccurrences(arr []int) bool {
	counter := make(map[int]int, len(arr)/2)
	for _, n := range arr {
		counter[n] += 1
	}

	seen := make(map[int]bool, len(counter))
	for _, v := range counter {
		if seen[v] {
			return false
		}

		seen[v] = true
	}

	return true
}
