package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(1)
func removeElements(head *helpers.ListNode, val int) *helpers.ListNode {
	if head == nil {
		return nil
	}

	dummy := &helpers.ListNode{Val: -1, Next: head}
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
func removeElements_Recursive(head *helpers.ListNode, val int) *helpers.ListNode {
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
