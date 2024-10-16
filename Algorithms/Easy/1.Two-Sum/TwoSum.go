package easy

// Time complexity O(n)
// Space complexity O(n)
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if _, ok := hashMap[diff]; ok {
			return []int{i, hashMap[diff]}
		}

		hashMap[nums[i]] = i
	}

	return nil
}
