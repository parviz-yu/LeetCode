package megium

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushBack(root)

	var isReversed bool
	for queue.Len() > 0 {
		level := make([]int, 0, queue.Len())
		for range queue.Len() {
			if !isReversed {
				root = queue.Remove(queue.Front()).(*TreeNode)
				level = append(level, root.Val)
				if root.Left != nil {
					queue.PushBack(root.Left)
				}
				if root.Right != nil {
					queue.PushBack(root.Right)
				}
			} else {
				root = queue.Remove(queue.Back()).(*TreeNode)
				level = append(level, root.Val)
				if root.Right != nil {
					queue.PushFront(root.Right)
				}
				if root.Left != nil {
					queue.PushFront(root.Left)
				}
			}
		}

		isReversed = !isReversed
		result = append(result, level)
	}

	return result
}
