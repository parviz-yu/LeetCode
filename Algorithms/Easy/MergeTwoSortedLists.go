package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(1)
func mergeTwoLists(list1 *helpers.ListNode, list2 *helpers.ListNode) *helpers.ListNode {
	dummy := &helpers.ListNode{}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	}

	if list2 != nil {
		curr.Next = list2
	}

	return dummy.Next
}
