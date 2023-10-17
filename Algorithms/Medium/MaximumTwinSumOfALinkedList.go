package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(1)
func pairSum(head *helpers.ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *helpers.ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	max_sum := 0
	for head != nil && prev != nil {
		max_sum = helpers.Max(max_sum, head.Val+prev.Val)
		head = head.Next
		prev = prev.Next
	}

	return max_sum
}
