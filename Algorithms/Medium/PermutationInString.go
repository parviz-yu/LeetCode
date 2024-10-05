package medium

// Time: O(n)
// Space: O(n)
func checkInclusion(s1 string, s2 string) bool {
	freq := make(map[byte]int)
	for i := range s1 {
		freq[s1[i]]++
	}

	var begin int
	for end := range s2 {
		char := s2[end]
		if _, ok := freq[char]; ok {
			freq[char]--
		}

		// shrink window
		for _, ok := freq[char]; !ok && begin <= end || freq[char] < 0; {
			if _, ok := freq[s2[begin]]; ok {
				freq[s2[begin]]++
			}

			begin++
		}

		if end-begin+1 == len(s1) {
			return true
		}
	}

	return false
}
