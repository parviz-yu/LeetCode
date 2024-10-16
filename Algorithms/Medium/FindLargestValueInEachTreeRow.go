package medium

import (
	"container/list"
	"math"

	"github.com/parviz-yu/LeetCode/Algorithms/helpers"
)

// Time: O(n)
// Space: O(n)
func largestValues(root *helpers.TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {

		maxVal := math.MinInt
		for range queue.Len() {
			node := queue.Remove(queue.Front()).(*helpers.TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			maxVal = max(maxVal, node.Val)
		}

		result = append(result, maxVal)
	}

	return result
}
