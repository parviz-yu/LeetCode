package easy

import (
	"slices"

	"github.com/parviz-yu/LeetCode/Algorithms/helpers"
)

// Time: O(n)
// Space: O(n)
func postorderTraversal_recursive(root *helpers.TreeNode) []int {
	return postorderDFS(root, []int{})
}

func postorderDFS(node *helpers.TreeNode, elems []int) []int {
	if node == nil {
		return elems
	}

	elems = postorderDFS(node.Left, elems)
	elems = postorderDFS(node.Right, elems)
	elems = append(elems, node.Val)

	return elems
}

// Time: O(n)
// Space: O(n)
func postorderTraversal(root *helpers.TreeNode) []int {
	var result []int
	var stack []*helpers.TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Right
		}
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		root = root.Left
	}

	slices.Reverse(result)
	return result
}
