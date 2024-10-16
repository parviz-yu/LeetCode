package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(nlogn)
// Space complexity: O(logn)
func sortList(head *ListNode) *ListNode {
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

func midList(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func mergeList(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
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
