package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(n)
// Space complexity: O(1)
func splitListToParts(head *ListNode, k int) []*ListNode {
	parts := make([]*ListNode, 0, k)

	curr, length := head, 0
	for curr != nil {
		length++
		curr = curr.Next
	}

	subLength := length / k
	reminders := length % k

	var prev *ListNode
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
