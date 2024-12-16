package medium

import "strings"

// Time: O(n)
// Space: O(n)
func reverseWords(s string) string {
	buf := make([]string, 0)
	var start, end int
	for end = range s {
		if s[end] == ' ' {
			if s[start:end] != "" {
				buf = append(buf, s[start:end])
			}
			start = end + 1
		}
	}
	if start <= end {
		buf = append(buf, s[start:])
	}

	for l, r := 0, len(buf)-1; l < r; {
		buf[l], buf[r] = buf[r], buf[l]
		l++
		r--
	}

	return strings.Join(buf, " ")
}

// Time: O(n)
// Space: O(n)
func reverseWords_lib(s string) string {
	buf := strings.Fields(s)
	for l, r := 0, len(buf)-1; l < r; l, r = l+1, r-1 {
		buf[l], buf[r] = buf[r], buf[l]
	}

	return strings.Join(buf, " ")
}
