package medium

// Time complexity: O(n)
// Space complexity: O(n)
func removeStars(s string) string {
	stack := make([]byte, 0, len(s))
	pop := func() {
		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			pop()
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
