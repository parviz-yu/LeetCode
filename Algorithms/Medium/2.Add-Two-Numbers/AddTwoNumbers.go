package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(N)
// Space complexity: O(1)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	addNext, val := 0, 0
	for l1 != nil || l2 != nil || addNext > 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		addNext, val = divMod(v1+v2+addNext, 10)
		curr.Next = &ListNode{val, nil}
		curr = curr.Next
	}

	return dummy.Next
}

func divMod(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}
