package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n)
// Space: O(1)
func deleteMiddle_1(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	middle := slow
	slow = head
	for slow.Next != middle {
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}

// Time: O(n)
// Space: O(1)
func deleteMiddle_2(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	slow := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = slow.Next.Next

	return head
}
