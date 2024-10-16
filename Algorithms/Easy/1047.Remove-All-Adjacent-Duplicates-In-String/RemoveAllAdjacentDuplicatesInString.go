package easy

// Time: O(n)
// Space: O(n)
func removeDuplicates(s string) string {
	stack := make([]rune, 0)

	for _, char := range s {
		if len(stack) == 0 || stack[len(stack)-1] != char {
			stack = append(stack, char)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}
