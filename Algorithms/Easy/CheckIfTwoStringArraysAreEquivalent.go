package easy

import "strings"

// Time: O(n)
// Time: O(n)
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	w1 := strings.Join(word1, "")
	w2 := strings.Join(word2, "")

	var i, j int
	for i < len(w1) && j < len(w2) {
		if w1[i] != w2[j] {
			return false
		}

		i++
		j++
	}

	return i == len(w1) && j == len(w2)
}
