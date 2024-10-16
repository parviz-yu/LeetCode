package easy

// Time complexity: O(n)
// Space complexity: O(1)
func longestCommonPrefix(strs []string) string {
	res := make([]byte, 0)
	for i := range strs[0] {
		for _, str := range strs {
			if i == len(str) || str[i] != strs[0][i] {
				return string(res)
			}
		}
		res = append(res, strs[0][i])
	}

	return string(res)
}
