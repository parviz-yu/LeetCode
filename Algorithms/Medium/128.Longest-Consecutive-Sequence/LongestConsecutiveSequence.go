package medium

// Time complexity: O(n)
// Space complexity: O(n)
func longestConsecutive(nums []int) int {
	hashSet := make(map[int]bool)
	longest_seq := 0

	for _, num := range nums {
		if !hashSet[num] {
			hashSet[num] = true
		}
	}

	for _, num := range nums {
		if !hashSet[num-1] {
			next := num + 1
			for hashSet[next] {
				next++
			}

			longest_seq = max(longest_seq, next-num)
		}
	}

	return longest_seq
}
