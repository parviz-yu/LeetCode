package easy

// Time: O(n)
// Space: O(n)
func clearDigits(s string) string {
	stack := make([]byte, 0)
	for i := range s {
		if '0' <= s[i] && s[i] <= '9' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
