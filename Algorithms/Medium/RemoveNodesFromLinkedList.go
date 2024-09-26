package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time: O(n)
// Space: O(1)
func removeNodes(head *helpers.ListNode) *helpers.ListNode {
	reverse := func(head *helpers.ListNode) *helpers.ListNode {
		var prev *helpers.ListNode
		for head != nil {
			next := head.Next
			head.Next = prev
			prev = head
			head = next
		}

		return prev
	}

	head = reverse(head)
	curr := head

	for curr.Next != nil {
		if curr.Val > curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return reverse(head)
}
