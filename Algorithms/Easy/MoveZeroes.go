package easy

// Time complexity: O(n)
// Space complexity: O(1)
func moveZeroes_1(nums []int) {
	lastNonZero := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[lastNonZero] = nums[lastNonZero], nums[i]
			lastNonZero++
		}
	}
}

// Time complexity: O(n)
// Space complexity: O(n)
func moveZeroes_2(nums []int) {
	numsCp := make([]int, len(nums))
	copy(numsCp, nums)
	l, r := 0, len(nums)-1
	for _, num := range numsCp {
		if num != 0 {
			nums[l] = num
			l++
		} else {
			nums[r] = num
			r--
		}
	}
}
