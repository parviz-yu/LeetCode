package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time: O(n)
// Space: O(n)
func invertTree_recursive(root *helpers.TreeNode) *helpers.TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree_recursive(root.Left)
	invertTree_recursive(root.Right)

	return root
}

// Time: O(n)
// Space: O(n)
func invertTree(root *helpers.TreeNode) *helpers.TreeNode {
	if root == nil {
		return nil
	}

	stack := []*helpers.TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return root
}
