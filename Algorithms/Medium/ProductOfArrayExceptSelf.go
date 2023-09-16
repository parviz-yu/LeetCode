package medium

// Time complexity: O(n)
// Space complexity: O(n)
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	prefix := make([]int, len(nums)+1)
	suffix := make([]int, len(nums)+1)

	for i := range prefix {
		prefix[i] = 1
		suffix[i] = 1
	}

	for i := range nums {
		prefix[i+1] = prefix[i] * nums[i]
	}

	for i := len(nums) - 1; i > 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}

	for i := 0; i < len(nums); i++ {
		result[i] = prefix[i] * suffix[i+1]
	}

	return result
}
