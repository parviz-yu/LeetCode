package easy

import "strings"

// Time complexity: O(n)
// Space complexity: O(n)
func wordPattern(pattern string, s string) bool {
	spl := strings.Split(s, " ")
	if len(spl) != len(pattern) {
		return false
	}

	spl2p := make(map[string]byte)
	p2spl := make(map[byte]string)

	for i := 0; i < len(spl); i++ {
		if v, ok := spl2p[spl[i]]; ok && v != pattern[i] {
			return false
		}

		if v, ok := p2spl[pattern[i]]; ok && v != spl[i] {
			return false
		}

		spl2p[spl[i]] = pattern[i]
		p2spl[pattern[i]] = spl[i]
	}

	return true
}
