package easy

// Time complexity - O(n)
// Space complexity - O(n)
func containsDuplicate(nums []int) bool {
	dict := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := dict[num]; ok {
			return true
		}

		dict[num] = struct{}{}
	}

	return false
}
