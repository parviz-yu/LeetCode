package easy

// First approach
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
		counter[t[i]]--
	}

	for k := range counter {
		if counter[k] != 0 {
			return false
		}
	}

	return true
}

// Second approach
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := [26]int{}
	for i := 0; i < len(s); i++ {
		ch1 := s[i] - 'a'
		ch2 := t[i] - 'a'

		chars[ch1]++
		chars[ch2]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}
