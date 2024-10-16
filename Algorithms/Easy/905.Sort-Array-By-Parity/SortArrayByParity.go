package easy

// Time: O(n)
// Space: O(1)
func sortArrayByParity(nums []int) []int {
	var k int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}

	return nums
}
