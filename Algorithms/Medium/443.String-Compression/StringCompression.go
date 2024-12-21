package medium

import "strconv"

// Time: O(n)
// Space: O(1)
func compress(chars []byte) int {
	var counter int
	var i int
	prev := chars[0]
	for j := range chars {
		curr := chars[j]
		if prev != curr {
			i += 1
			if counter > 1 {
				counterStr := strconv.Itoa(counter)
				for k := range counterStr {
					chars[i] = counterStr[k]
					i++
				}
			}

			chars[i] = curr
			prev = curr
			counter = 0
		}

		counter++
	}
	i++

	if counter > 1 {
		counterStr := strconv.Itoa(counter)
		for k := range counterStr {
			chars[i] = counterStr[k]
			i++
		}
	}

	return i
}
