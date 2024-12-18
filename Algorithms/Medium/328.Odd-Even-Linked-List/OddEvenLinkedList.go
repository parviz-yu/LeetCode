package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n)
// Space: O(1)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	oddList := head
	evenList := head.Next

	oddPtr := oddList
	evenPtr := evenList
	for oddPtr != nil && oddPtr.Next != nil && evenPtr != nil && evenPtr.Next != nil {
		oddPtr.Next = oddPtr.Next.Next
		evenPtr.Next = evenPtr.Next.Next
		oddPtr = oddPtr.Next
		evenPtr = evenPtr.Next
	}

	if evenPtr != nil {
		evenPtr.Next = nil
	}

	oddPtr.Next = evenList
	return oddList
}
