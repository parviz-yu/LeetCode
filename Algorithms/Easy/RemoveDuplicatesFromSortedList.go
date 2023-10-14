package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(1)
func deleteDuplicates(head *helpers.ListNode) *helpers.ListNode {
	dummy := &helpers.ListNode{Val: -101, Next: head}
	curr := dummy
	for curr.Next != nil {
		if curr.Val == head.Val {
			curr.Next = head.Next
		} else {
			curr = curr.Next
		}
		head = head.Next
	}

	return dummy.Next
}
