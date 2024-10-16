package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(n)
// Space complexity: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -101, Next: head}
	curr := dummy
	for curr.Next != nil {
		if curr.Val == head.Val {
			curr.Next = head.Next
		} else {
			curr = curr.Next
		}
		head = head.Next
	}

	return dummy.Next
}
