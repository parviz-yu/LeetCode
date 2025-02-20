package medium

// Time: O(2^n)
// Space: O(2^n)
func findDifferentBinaryString(nums []string) string {
	seen := make(map[string]bool)
	for _, num := range nums {
		seen[num] = true
	}

	var answer string
	var found bool
	var backtrack func(acc []byte)
	backtrack = func(acc []byte) {
		if found {
			return
		}

		if len(acc) == len(nums) {
			if !seen[string(acc)] {
				answer = string(acc)
				found = true
			}

			return
		}

		for _, char := range []byte("01") {
			acc = append(acc, char)
			backtrack(acc)
			acc = acc[:len(acc)-1]
		}
	}

	backtrack([]byte{})
	return answer
}
