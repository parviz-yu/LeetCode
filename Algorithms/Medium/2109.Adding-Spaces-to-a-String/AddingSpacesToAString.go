package medium

// Time: O(n)
// Space: O(n)
func addSpaces(s string, spaces []int) string {
	buf := make([]byte, 0, len(spaces)+1)
	var start int
	for _, end := range spaces {
		buf = append(buf, s[start:end]...)
		buf = append(buf, []byte(" ")...)
		start = end
	}

	buf = append(buf, s[start:]...)
	return string(buf)
}
