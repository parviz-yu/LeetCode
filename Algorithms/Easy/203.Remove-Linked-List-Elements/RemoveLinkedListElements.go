package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(n)
// Space complexity: O(1)
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Val: -1, Next: head}
	curr := dummy
	for curr != nil && curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return dummy.Next
}

// Time complexity: O(n)
// Space complexity: O(1)
func removeElements_Recursive(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	second := removeElements(head.Next, val)
	if head.Val == val {
		return second
	} else {
		head.Next = second
		return head
	}
}
