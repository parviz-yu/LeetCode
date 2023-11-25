package easy

// Time complexity: O(n)
// Space complexity: O(n)
func findDifference(nums1 []int, nums2 []int) [][]int {
	answer := make([][]int, 2)
	nums1set := arrayToSet(nums1)
	nums2set := arrayToSet(nums2)

	for num := range nums1set {
		if _, ok := nums2set[num]; !ok {
			answer[0] = append(answer[0], num)
		}
	}

	for num := range nums2set {
		if _, ok := nums1set[num]; !ok {
			answer[1] = append(answer[1], num)
		}
	}

	return answer
}

func arrayToSet(array []int) map[int]struct{} {
	set := make(map[int]struct{})
	for _, num := range array {
		set[num] = struct{}{}
	}

	return set
}
