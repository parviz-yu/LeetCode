package medium

// Time: O(n)
// Space: O(n)
func removeOccurrences(s string, part string) string {
	stack := make([]byte, 0)
	n := len(part)
	for i := range s {
		stack = append(stack, s[i])
		if len(stack) >= n && string(stack[len(stack)-n:]) == part {
			stack = stack[:len(stack)-n]
		}
	}

	return string(stack)
}
