package easy

import "strings"

// Time complexity: O(n)
// Space complexity: O(1)
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	s = strings.ToLower(s)

	for l < r {
		if !isAlphaNum(s[l]) {
			l++
		} else if !isAlphaNum(s[r]) {
			r--
		} else {
			if s[l] != s[r] {
				return false
			}

			l++
			r--
		}
	}

	return true
}

func isAlphaNum(c byte) bool {
	if c >= 48 && c <= 57 || c >= 97 && c <= 122 {
		return true
	}

	return false
}
