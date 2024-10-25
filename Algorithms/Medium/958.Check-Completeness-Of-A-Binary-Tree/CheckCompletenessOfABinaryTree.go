package medium

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func isCompleteTree(root *TreeNode) bool {
	queue := list.New()
	queue.PushBack(root)

	var metGap bool
	for queue.Len() > 0 {

		for range queue.Len() {
			root = queue.Remove(queue.Front()).(*TreeNode)
			if root != nil && metGap {
				return false
			}
			if root != nil {
				queue.PushBack(root.Left)
				queue.PushBack(root.Right)
			}
			if root == nil && !metGap {
				metGap = true
			}
		}
	}

	return true
}
