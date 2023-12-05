package easy

// Time complexity: O(n)
// Space complexity: O(n)
func removeDuplicates_1(nums []int) int {
	seen := make(map[int]struct{})
	idx := 0
	for _, num := range nums {
		if _, ok := seen[num]; !ok {
			nums[idx] = num
			idx++
			seen[num] = struct{}{}
		}
	}

	return len(seen)
}

// Time complexity: O(n)
// Space complexity: O(1)
func removeDuplicates_2(nums []int) int {
	idx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[idx] = nums[i]
			idx++
		}
	}

	return idx
}
