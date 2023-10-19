package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(1)
func swapNodes(head *helpers.ListNode, k int) *helpers.ListNode {
	first := head
	for i := 0; i < k-1; i++ {
		first = first.Next
	}

	curr, second := first, head
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}

	curr.Val, second.Val = second.Val, curr.Val
	return head
}
