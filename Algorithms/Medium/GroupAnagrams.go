package medium

import (
	"sort"
	"strings"
)

// Time complexity: O(nklog(k))
//   k - length of string
//   n - length of array
// Space complexity: O(n)
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	hashMap := make(map[string][]string)

	for _, str := range strs {
		sorted := sortWord(str)
		hashMap[sorted] = append(hashMap[sorted], str)
	}

	for _, v := range hashMap {
		result = append(result, v)
	}

	return result
}

func sortWord(word string) string {
	sorted := strings.Split(word, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}
