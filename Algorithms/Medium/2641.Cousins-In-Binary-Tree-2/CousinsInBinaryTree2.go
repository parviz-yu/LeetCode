package medium

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func replaceValueInTree(root *TreeNode) *TreeNode {
	queue := list.New()
	root.Val = 0
	queue.PushBack(root)

	for queue.Len() > 0 {
		pairs := make(map[*TreeNode][]*TreeNode)
		var levelSum int
		for range queue.Len() {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				pairs[node] = append(pairs[node], node.Left)
				levelSum += node.Left.Val
			}
			if node.Right != nil {
				pairs[node] = append(pairs[node], node.Right)
				levelSum += node.Right.Val
			}
		}

		for _, childs := range pairs {

			// O(1): child len <= 2
			var currSum int
			for _, child := range childs {
				currSum += child.Val
			}

			// O(1): child len <= 2
			for _, child := range childs {
				child.Val = levelSum - currSum
				queue.PushBack(child)
			}
		}
	}

	return root
}
