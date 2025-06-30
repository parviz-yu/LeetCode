package easy

// Time: O(n)
// Space: O(1)
func maxPower(s string) int {
	char := s[0]
	currLength := 0
	maxLength := 1

	for i := range s {
		if s[i] == char {
			currLength++
		} else {
			maxLength = max(maxLength, currLength)
			currLength = 1
			char = s[i]
		}
	}

	maxLength = max(maxLength, currLength)
	return maxLength
}
