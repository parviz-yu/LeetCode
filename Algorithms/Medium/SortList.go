package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(nlogn)
// Space complexity: O(logn)
func sortList(head *helpers.ListNode) *helpers.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left := head
	right := midList(head)
	tmp := right.Next
	right.Next = nil
	right = tmp

	left = sortList(left)
	right = sortList(right)
	return mergeList(left, right)
}

func midList(head *helpers.ListNode) *helpers.ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func mergeList(left, right *helpers.ListNode) *helpers.ListNode {
	dummy := &helpers.ListNode{}
	curr := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			curr.Next = left
			left = left.Next
		} else {
			curr.Next = right
			right = right.Next
		}
		curr = curr.Next
	}

	if left != nil {
		curr.Next = left
	}
	if right != nil {
		curr.Next = right
	}

	return dummy.Next
}
