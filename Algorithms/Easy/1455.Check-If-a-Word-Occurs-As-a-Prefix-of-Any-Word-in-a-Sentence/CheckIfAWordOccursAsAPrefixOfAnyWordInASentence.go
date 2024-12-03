package easy

import "strings"

// Time: O(m*n)
// Space: O(n)
func isPrefixOfWord(sentence string, searchWord string) int {
	s := strings.Split(sentence, " ")
	for i, w := range s {
		if len(w) >= len(searchWord) && hasPrefix(w, searchWord) {
			return i + 1
		}
	}

	return -1
}

func hasPrefix(s, prefix string) bool {
	for i := range prefix {
		if s[i] != prefix[i] {
			return false
		}
	}

	return true
}
