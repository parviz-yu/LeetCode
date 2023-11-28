package easy

// Time complexity: O(n)
// Space complexity: O(1)
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return checkRemain(s, l+1, r) || checkRemain(s, l, r-1)
		}
		l++
		r--
	}

	return true
}

func checkRemain(s string, l int, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
