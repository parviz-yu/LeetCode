package easy

// Time complexity: O(n)
// Space complexity: O(n)
func findDisappearedNumbers_1(nums []int) []int {
	counter := make([]int, len(nums)+1)
	for _, num := range nums {
		counter[num]++
	}

	ans := make([]int, 0)
	for i := 1; i < len(counter); i++ {
		if counter[i] == 0 {
			ans = append(ans, i)
		}
	}

	return ans
}

// Time complexity: O(n)
// Space complexity: O(1)
func findDisappearedNumbers_2(nums []int) []int {
	for _, num := range nums {
		// absolute value
		if num < 0 {
			num *= -1
		}

		idx := num - 1
		if nums[idx] > 0 {
			nums[idx] *= -1
		}
	}

	ans := make([]int, 0)
	for i := range nums {
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}

	return ans
}
