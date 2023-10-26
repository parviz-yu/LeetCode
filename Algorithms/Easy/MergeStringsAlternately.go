package easy

// Time complexity: O(n)
// Space complexity: O(1)
func mergeAlternately(word1 string, word2 string) string {
	m, n := len(word1), len(word2)
	res := make([]byte, 0, m+n)

	ptr1, ptr2 := 0, 0
	for ptr1 < m && ptr2 < n {
		if ptr1 == ptr2 {
			res = append(res, word1[ptr1])
			ptr1++
		} else {
			res = append(res, word2[ptr2])
			ptr2++
		}
	}

	if ptr1 < m {
		res = append(res, word1[ptr1:]...)
	}
	if ptr2 < n {
		res = append(res, word2[ptr2:]...)
	}

	return string(res)
}
