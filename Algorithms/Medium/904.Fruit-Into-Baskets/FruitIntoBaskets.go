package medium

// Time: O(n)
// Space: O(1), len(map) = 2
func totalFruit(fruits []int) int {
	var begin, maxNum int
	basket := make(map[int]int)

	for end := range fruits {
		_, found := basket[fruits[end]]
		for !found && len(basket) == 2 {
			basket[fruits[begin]]--
			if basket[fruits[begin]] == 0 {
				delete(basket, fruits[begin])
			}

			begin++
		}

		basket[fruits[end]]++
		maxNum = max(maxNum, end-begin+1)
	}

	return maxNum
}
