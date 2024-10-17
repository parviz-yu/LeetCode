package medium

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func deepestLeavesSum(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)

	var levelSum int
	for queue.Len() > 0 {

		levelSum = 0
		for range queue.Len() {
			node := queue.Remove(queue.Front()).(*TreeNode)
			levelSum += node.Val

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return levelSum
}
