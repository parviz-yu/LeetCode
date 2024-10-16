package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(n)
// Space complexity: O(1)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		nxt := head.Next
		head.Next = prev
		prev = head
		head = nxt
	}

	return prev
}
