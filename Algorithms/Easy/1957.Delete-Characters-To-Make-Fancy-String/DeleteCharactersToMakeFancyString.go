package easy

// Time: O(n)
// Space: O(n)
func makeFancyString(s string) string {
	var stack []byte
	var counter int
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			if (top == s[i] && counter < 2) || top != s[i] {
				stack = append(stack, s[i])
			}
			if top != s[i] {
				counter = 0
			}
		} else {
			stack = append(stack, s[i])
		}

		counter++
	}

	return string(stack)
}
