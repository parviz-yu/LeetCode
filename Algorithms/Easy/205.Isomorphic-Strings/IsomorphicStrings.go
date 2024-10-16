package easy

// Time complexity: O(n)
// Space complexity: O(n)
func isIsomorphic(s string, t string) bool {
	mapST := make(map[byte]byte)
	mapTS := make(map[byte]byte)

	for i := range s {
		if _, ok := mapST[s[i]]; ok && mapST[s[i]] != t[i] {
			return false
		}

		if _, ok := mapTS[t[i]]; ok && mapTS[t[i]] != s[i] {
			return false
		}

		mapST[s[i]] = t[i]
		mapTS[t[i]] = s[i]
	}

	return true
}
