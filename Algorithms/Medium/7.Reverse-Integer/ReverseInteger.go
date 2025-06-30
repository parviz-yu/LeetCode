package medium

const (
	MaxInt32 = 1<<31 - 1
	MinInt32 = -1 << 31
)

// Time: O(n)
// Space: O(1)
func reverse(x int) int {
	var result int
	var isNegative bool
	if x < 0 {
		isNegative = true
		x *= -1
	}

	for x != 0 {
		result = (result * 10) + (x % 10)
		x /= 10
	}

	if isNegative {
		result *= -1
	}

	if !(result >= MinInt32 && result <= MaxInt32) {
		return 0
	}

	return result
}
