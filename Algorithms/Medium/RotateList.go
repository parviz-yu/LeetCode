package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(1)
func rotateRight(head *helpers.ListNode, k int) *helpers.ListNode {
	if head == nil {
		return head
	}

	len, curr := 1, head
	for curr.Next != nil {
		curr = curr.Next
		len++
	}

	curr.Next = head
	k = k % len
	for i := 0; i < len-k-1; i++ {
		head = head.Next
	}

	curr = head.Next
	head.Next = nil

	return curr
}
