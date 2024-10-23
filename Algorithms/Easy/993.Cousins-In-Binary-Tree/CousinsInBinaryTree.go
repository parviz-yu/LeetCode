package easy

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func isCousins(root *TreeNode, x int, y int) bool {
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		var res int
		for range queue.Len() {
			root = queue.Remove(queue.Front()).(*TreeNode)
			if root != nil {
				l := root.Left
				r := root.Right

				if l != nil && (l.Val == x || l.Val == y) {
					res++
				} else if r != nil && (r.Val == x || r.Val == y) {
					res++
				}

				queue.PushBack(l)
				queue.PushBack(r)
			}
		}
		if res == 2 {
			return true
		}
	}

	return false
}
