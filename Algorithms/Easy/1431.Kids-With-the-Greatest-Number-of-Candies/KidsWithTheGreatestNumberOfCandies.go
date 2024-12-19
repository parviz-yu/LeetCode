package easy

// Time: O(n)
// Space: O(1)
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxVal := candies[0]
	for i := range candies {
		maxVal = max(maxVal, candies[i])
	}

	result := make([]bool, len(candies))
	for i := range candies {
		if candies[i]+extraCandies >= maxVal {
			result[i] = true
		}
	}

	return result
}
