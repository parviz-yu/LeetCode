package hard

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(nk)
// Space: O(1)
func mergeKLists_1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	sorted := lists[0]
	for i := 1; i < len(lists); i++ {
		sorted = merge(sorted, lists[i])
	}

	return sorted
}

// Time: O(nlogk)
// Space: O(n)
func mergeKLists_2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {

		sorted := make([]*ListNode, 0)
		for i := 0; i < len(lists); i += 2 {
			l1 := lists[i]
			var l2 *ListNode
			if i+1 < len(lists) {
				l2 = lists[i+1]
			}

			sorted = append(sorted, merge(l1, l2))
		}
		lists = sorted
	}

	return lists[0]
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}

		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}

	if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}
