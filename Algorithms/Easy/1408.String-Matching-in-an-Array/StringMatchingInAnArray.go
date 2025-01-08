package easy

import (
	"sort"
	"strings"
)

// Time: O(n^2)
// Space: O(n)
func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	var result []string
	seen := make(map[string]bool)
	for r := len(words) - 1; r >= 0; r-- {
		for l := 0; l < r; l++ {
			if strings.Contains(words[r], words[l]) && !seen[words[l]] {
				result = append(result, words[l])
				seen[words[l]] = true
			}
		}
	}

	return result
}
