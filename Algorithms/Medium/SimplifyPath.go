package medium

import "strings"

// Time: O(n)
// Space: O(n)
func simplifyPath(path string) string {
	stack := make([]string, 0)
	pathSplit := strings.Split(path, "/")

	for _, word := range pathSplit {
		if word == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if word != "" && word != "." {
			stack = append(stack, word)
		}
	}

	return "/" + strings.Join(stack, "/")
}
