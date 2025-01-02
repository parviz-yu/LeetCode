package medium

// Time: O(n)
// Space: O(n)
func vowelStrings(words []string, queries [][]int) []int {
	prefix := make([]int, len(words)+1)
	for i := range len(words) {
		word := words[i]
		if isVowel(word[0]) && isVowel(word[len(word)-1]) {
			prefix[i+1] = prefix[i] + 1
		} else {
			prefix[i+1] = prefix[i]
		}
	}

	count := make([]int, len(queries))
	for i := range len(queries) {
		l, r := queries[i][0], queries[i][1]
		count[i] = prefix[r+1] - prefix[l]
	}

	return count
}

func isVowel(char byte) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}
