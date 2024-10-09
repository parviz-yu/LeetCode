package easy

// Time: O(m + n)
// Space: O(n)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreater := make(map[int]int, 0)
	stack := make([]int, 0)

	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGreater[top] = num
		}

		stack = append(stack, num)
	}

	for i, num := range nums1 {
		if val, ok := nextGreater[num]; ok {
			nums1[i] = val
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}
