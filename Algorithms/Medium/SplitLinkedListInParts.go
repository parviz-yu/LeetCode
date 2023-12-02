package medium

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time complexity: O(n)
// Space complexity: O(1)
func splitListToParts(head *helpers.ListNode, k int) []*helpers.ListNode {
	parts := make([]*helpers.ListNode, 0, k)

	curr, length := head, 0
	for curr != nil {
		length++
		curr = curr.Next
	}

	subLength := length / k
	reminders := length % k

	var prev *helpers.ListNode
	curr = head
	for ; k > 0; k-- {
		parts = append(parts, curr)

		skip := subLength
		if reminders > 0 {
			skip++
			reminders--
		}

		for ; skip > 0; skip-- {
			prev = curr
			curr = curr.Next
		}

		if prev != nil {
			prev.Next = nil
		}
	}

	return parts
}
