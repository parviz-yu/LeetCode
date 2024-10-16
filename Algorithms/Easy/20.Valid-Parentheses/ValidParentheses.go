package easy

// Time complexity: O(N)
// Space complexity: O(N)
func isValid(s string) bool {
	stack := make([]byte, 0, len(s)/2)
	for _, c := range []byte(s) {
		if c == '{' || c == '[' || c == '(' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}

			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			pair := string([]byte{popped, c})
			if !contains([]string{"()", "{}", "[]"}, pair) {
				return false
			}
		}
	}

	return len(stack) == 0
}

func contains(valid []string, pair string) bool {
	for _, p := range valid {
		if p == pair {
			return true
		}
	}

	return false
}
