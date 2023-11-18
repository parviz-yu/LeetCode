package helpers

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Sum(nums []int) int {
	total := 0
	for _, val := range nums {
		total += val
	}

	return total
}
