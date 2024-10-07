package easy

// Time: O(n)
// Space: O(n)
func minLength(s string) int {
	res := make([]rune, 0)

	for _, char := range s {
		if len(res) > 0 {
			if char == 'B' && res[len(res)-1] == 'A' {
				res = res[:len(res)-1]
				continue
			} else if char == 'D' && res[len(res)-1] == 'C' {
				res = res[:len(res)-1]
				continue
			}
		}

		res = append(res, char)
	}

	return len(res)
}
