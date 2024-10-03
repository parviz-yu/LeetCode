package easy

import "unicode"

// Time: O(n)
// Space: O(n)
func makeGood(s string) string {
	goodString := make([]rune, 0)

	for _, char := range s {
		if len(goodString) > 0 {
			elem := goodString[len(goodString)-1]
			if unicode.ToLower(elem) == unicode.ToLower(char) && elem != char {
				goodString = goodString[:len(goodString)-1]
				continue
			}
		}

		goodString = append(goodString, char)
	}

	return string(goodString)
}
