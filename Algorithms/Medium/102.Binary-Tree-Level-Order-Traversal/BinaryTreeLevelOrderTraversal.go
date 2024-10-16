package medium

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)

	var result [][]int
	for queue.Len() > 0 {

		var levelElems []int
		for range queue.Len() {
			node := queue.Remove(queue.Front()).(*TreeNode)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}

			levelElems = append(levelElems, node.Val)
		}

		result = append(result, levelElems)
	}

	return result
}
