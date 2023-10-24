package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(1)
func swapPairs(head *helpers.ListNode) *helpers.ListNode {
	dummy := &helpers.ListNode{Val: 0, Next: head}
	prev, curr := dummy, head

	for curr != nil && curr.Next != nil {
		next_next := curr.Next.Next
		next := curr.Next

		curr.Next = next_next
		next.Next = curr
		prev.Next = next

		prev = curr
		curr = next_next
	}

	return dummy.Next
}
