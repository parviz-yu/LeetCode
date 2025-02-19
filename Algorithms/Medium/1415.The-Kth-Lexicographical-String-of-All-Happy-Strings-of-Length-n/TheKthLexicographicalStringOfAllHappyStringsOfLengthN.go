package medium

import "sort"

// Time: O(2^n)
// Space: O(n)
func getHappyString(n int, k int) string {
	alphabet := []byte("abc")

	result := make([]string, 0, n)
	var backtrack func(acc []byte, prevChar byte)
	backtrack = func(acc []byte, prevChar byte) {
		if len(acc) == n {
			result = append(result, string(acc))
			return
		}

		for _, char := range alphabet {
			if char != prevChar {
				acc = append(acc, char)
				backtrack(acc, char)
				acc = acc[:len(acc)-1]
			}
		}
	}

	backtrack([]byte{}, '0')

	if k > len(result) {
		return ""
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result[k-1]
}
