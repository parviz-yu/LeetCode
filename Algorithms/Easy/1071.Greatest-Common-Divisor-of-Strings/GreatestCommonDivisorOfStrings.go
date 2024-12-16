package easy

import "strings"

// Time: O(n)
// Space: O(n)
func reverseVowels(s string) string {
	const vowels string = "aAeEiIoOuU"

	result := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		cL, cR := string(s[l]), string(s[r])
		if strings.Contains(vowels, cR) && strings.Contains(vowels, cL) {
			result[l], result[r] = result[r], result[l]
			l++
			r--
		}
		if !strings.Contains(vowels, cL) {
			l++
		}
		if !strings.Contains(vowels, cR) {
			r--
		}
	}

	return string(result)
}
