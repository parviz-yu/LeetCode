package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(n)
// Space complexity: O(1)
func rotateRight(head *ListNode, k int) *ListNode {
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
