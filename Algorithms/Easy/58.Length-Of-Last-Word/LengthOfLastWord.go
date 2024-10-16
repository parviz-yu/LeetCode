package easy

// Time complexity: O(n)
// Space complexity: O(1)
func lengthOfLastWord(s string) int {
	i, length := len(s)-1, 0
	for s[i] == ' ' {
		i--
	}

	for i > -1 && s[i] != ' ' {
		length++
		i--
	}

	return length
}
