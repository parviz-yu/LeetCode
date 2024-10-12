package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time: O(n)
// Space: O(n)
func hasPathSum_recursive(root *helpers.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	leftChild := hasPathSum_recursive(root.Left, targetSum-root.Val)
	rightChild := hasPathSum_recursive(root.Right, targetSum-root.Val)
	return leftChild || rightChild
}

// Time: O(n)
// Space: O(n)
func hasPathSum(root *helpers.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	type tuple struct {
		node    *helpers.TreeNode
		currSum int
	}

	stack := []tuple{tuple{root, targetSum}}
	pop := func() (*helpers.TreeNode, int) {
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return pair.node, pair.currSum
	}

	for len(stack) > 0 {
		node, currSum := pop()
		if currSum == node.Val && node.Left == nil && node.Right == nil {
			return true
		}

		if node.Left != nil {
			stack = append(stack, tuple{node.Left, currSum - node.Val})
		}
		if node.Right != nil {
			stack = append(stack, tuple{node.Right, currSum - node.Val})
		}
	}

	return false
}
