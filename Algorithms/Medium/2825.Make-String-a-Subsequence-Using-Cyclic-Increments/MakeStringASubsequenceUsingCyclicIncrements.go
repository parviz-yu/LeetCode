package medium

// Time: O(n)
// Space: O(1)
func canMakeSubsequence(str1 string, str2 string) bool {
	if len(str2) > len(str1) {
		return false
	}

	var j int
	for i := 0; i < len(str1) && j < len(str2); i++ {
		var nextStr byte
		if str1[i] == 'z' {
			nextStr = 'a'
		} else {
			nextStr = str1[i] + 1
		}

		if str1[i] == str2[j] || nextStr == str2[j] {
			j++
		}
	}

	return j == len(str2)
}
