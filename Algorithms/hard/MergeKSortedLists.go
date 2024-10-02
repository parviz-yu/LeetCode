package hard

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time: O(nk)
// Space: O(1)
func mergeKLists_1(lists []*helpers.ListNode) *helpers.ListNode {
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
func mergeKLists_2(lists []*helpers.ListNode) *helpers.ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {

		sorted := make([]*helpers.ListNode, 0)
		for i := 0; i < len(lists); i += 2 {
			l1 := lists[i]
			var l2 *helpers.ListNode
			if i+1 < len(lists) {
				l2 = lists[i+1]
			}

			sorted = append(sorted, merge(l1, l2))
		}
		lists = sorted
	}

	return lists[0]
}

func merge(l1, l2 *helpers.ListNode) *helpers.ListNode {
	dummy := &helpers.ListNode{}
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
