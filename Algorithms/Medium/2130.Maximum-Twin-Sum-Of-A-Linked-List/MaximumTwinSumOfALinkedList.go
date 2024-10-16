package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(n)
// Space complexity: O(1)
func pairSum(head *ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	max_sum := 0
	for head != nil && prev != nil {
		max_sum = max(max_sum, head.Val+prev.Val)
		head = head.Next
		prev = prev.Next
	}

	return max_sum
}
