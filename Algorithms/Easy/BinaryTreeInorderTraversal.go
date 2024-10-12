package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time: O(n)
// Space: O(n)
func inorderTraversal_recursive(root *helpers.TreeNode) []int {
	result := make([]int, 0)
	return inorderDFS(root, result)
}

func inorderDFS(node *helpers.TreeNode, elems []int) []int {
	if node == nil {
		return elems
	}

	elems = inorderDFS(node.Left, elems)
	elems = append(elems, node.Val)
	elems = inorderDFS(node.Right, elems)
	return elems
}

// Time: O(n)
// Space: O(n)
func inorderTraversal(root *helpers.TreeNode) []int {
	var result []int
	var stack []*helpers.TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		result = append(result, root.Val)
		root = root.Right
	}

	return result
}
