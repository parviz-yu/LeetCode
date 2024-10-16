package easy

// Time complexity: O(m * n)
// Space complexity: O(m)
func numUniqueEmails(emails []string) int {
	seen := make(map[string]struct{})
	for _, email := range emails {
		idx := atSignIdx(email)
		validEmail := make([]byte, 0, len(email))
		for i := 0; i < len(email); i++ {
			if email[i] == '+' || email[i] == '@' {
				break
			}
			if email[i] == '.' {
				continue
			}
			validEmail = append(validEmail, email[i])
		}

		validEmail = append(validEmail, email[idx:]...)
		seen[string(validEmail)] = struct{}{}
	}

	return len(seen)
}

func atSignIdx(email string) int {
	for i := 0; i < len(email); i++ {
		if email[i] == '@' {
			return i
		}
	}

	return -1
}
