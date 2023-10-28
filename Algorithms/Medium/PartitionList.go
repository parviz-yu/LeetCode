package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(1)
func partition(head *helpers.ListNode, x int) *helpers.ListNode {
	less, more := helpers.ListNode{}, helpers.ListNode{}
	l_ptr, m_ptr := &less, &more

	for head != nil {
		if head.Val < x {
			l_ptr.Next = head
			l_ptr = l_ptr.Next
		} else {
			m_ptr.Next = head
			m_ptr = m_ptr.Next
		}

		head = head.Next
	}

	m_ptr.Next = nil
	l_ptr.Next = more.Next
	return less.Next
}
