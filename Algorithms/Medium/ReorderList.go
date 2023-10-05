package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(n)
func reorderList_1(head *helpers.ListNode) {
	slow, fast := head, head
	stack := make([]*helpers.ListNode, 0)
	pop := func() *helpers.ListNode {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return top
	}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	end := slow
	for slow != nil {
		stack = append(stack, slow)
		slow = slow.Next
	}

	for head != end {
		nxt := head.Next
		head.Next = pop()
		head.Next.Next = nxt
		head = nxt
	}

	end.Next = nil
}

// Time complexity: O(n)
// Space complexity: O(1)
func reorderList_2(head *helpers.ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *helpers.ListNode
	curr := slow.Next
	slow.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	first, curr := head, prev
	for curr != nil {
		tmp1, tmp2 := first.Next, curr.Next
		first.Next = curr
		curr.Next = tmp1
		first, curr = tmp1, tmp2
	}
}
