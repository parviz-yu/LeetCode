package easy

// Time: O(nm)
// Space: O(1)
func firstPalindrome(words []string) string {
	isPalindrome := func(s string) bool {
		l, r := 0, len(s)-1
		for l < r {
			if s[l] != s[r] {
				return false
			}

			l++
			r--
		}

		return true
	}

	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}

	return ""
}
