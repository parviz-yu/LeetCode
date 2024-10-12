package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time: O(n)
// Space: O(n)
func preorderTraversal_recursive(root *helpers.TreeNode) []int {
	return preorderDFS(root, []int{})
}

func preorderDFS(node *helpers.TreeNode, elems []int) []int {
	if node == nil {
		return elems
	}

	elems = append(elems, node.Val)
	elems = preorderDFS(node.Left, elems)
	elems = preorderDFS(node.Right, elems)
	return elems
}

// Time: O(n)
// Space: O(n)
func preorderTraversal(root *helpers.TreeNode) []int {
	var result []int
	var stack []*helpers.TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		root = root.Right
	}

	return result
}
