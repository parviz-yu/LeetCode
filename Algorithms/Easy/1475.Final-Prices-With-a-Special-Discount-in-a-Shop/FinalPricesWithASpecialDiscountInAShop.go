package easy

// Time: O(n)
// Space: O(n)
func finalPrices(prices []int) []int {
	n := len(prices)
	stack := make([]int, 0) // Monotonically Increasing Stack
	for i := n - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] > prices[i] {
			stack = stack[:len(stack)-1]
		}

		previousPrice := prices[i]
		if len(stack) != 0 {
			prices[i] -= stack[len(stack)-1]
		}

		stack = append(stack, previousPrice)
	}

	return prices
}
