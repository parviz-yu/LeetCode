package medium

import (
	"container/list"

	"github.com/parviz-yu/LeetCode/Algorithms/helpers"
)

// Time: O(n)
// Space: O(n)
func levelOrder(root *helpers.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)

	var result [][]int
	for queue.Len() > 0 {

		var levelElems []int
		for range queue.Len() {
			node := queue.Remove(queue.Front()).(*helpers.TreeNode)

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
