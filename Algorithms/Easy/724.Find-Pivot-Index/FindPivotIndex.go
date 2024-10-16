package easy

// Time complexity: O(n)
// Space complexity: O(1)
func pivotIndex(nums []int) int {
	leftSum, rightSum := 0, sum(nums)
	for i, val := range nums {
		rightSum -= val
		if leftSum == rightSum {
			return i
		} else {
			leftSum += val
		}
	}

	return -1
}

func sum(nums []int) int {
	total := 0
	for _, val := range nums {
		total += val
	}

	return total
}
