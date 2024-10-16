package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n)
// Space: O(1)
func removeNodes(head *ListNode) *ListNode {
	reverse := func(head *ListNode) *ListNode {
		var prev *ListNode
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
