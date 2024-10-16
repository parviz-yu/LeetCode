package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(n)
// Space complexity: O(1)
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
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
