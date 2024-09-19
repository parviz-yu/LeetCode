package medium

// Time: O(n)
// Space: O(1)
func removeDuplicates(nums []int) int {
	var k int
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[k] {
			count++
		} else {
			count = 1
		}
		if count < 3 {
			k++
			nums[k] = nums[i]
		}
	}

	return k + 1
}
