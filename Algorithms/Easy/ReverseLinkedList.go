package easy

import (
	"github.com/parviz-yu/LeetCode/Algorithms/helpers"
)

// Time complexity: O(n)
// Space complexity: O(1)
func reverseList(head *helpers.ListNode) *helpers.ListNode {
	var prev *helpers.ListNode

	for head != nil {
		nxt := head.Next
		head.Next = prev
		prev = head
		head = nxt
	}

	return prev
}
