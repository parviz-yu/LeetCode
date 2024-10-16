package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(n)
// Space complexity: O(1)
func partition(head *ListNode, x int) *ListNode {
	less, more := ListNode{}, ListNode{}
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
