package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time: O(n)
// Space: O(1)
func mergeInBetween(list1 *helpers.ListNode, a int, b int, list2 *helpers.ListNode) *helpers.ListNode {
	var i int
	curr := list1
	for ; i < a-1; i++ {
		curr = curr.Next
	}

	part1 := curr

	for ; i < b; i++ {
		curr = curr.Next
	}

	part2 := curr.Next

	part1.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}

	list2.Next = part2

	return list1
}
