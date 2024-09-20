package medium

// Time: O(n)
// Space: O(1)
func maxVowels(s string, k int) int {
	var begin, currNum, maxNum int
	for end := range s {
		if isVowel(string(s[end])) {
			currNum++
			maxNum = max(maxNum, currNum)
		}

		if end-begin+1 == k {
			if isVowel(string(s[begin])) {
				currNum--
			}
			begin++
		}
	}

	return maxNum
}

func isVowel(char string) bool {
	switch char {
	case "a", "e", "i", "o", "u":
		return true
	default:
		return false
	}
}
