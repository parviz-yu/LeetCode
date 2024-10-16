package easy

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func postorderTraversal_recursive(root *TreeNode) []int {
	return postorderDFS(root, []int{})
}

func postorderDFS(node *TreeNode, elems []int) []int {
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
func postorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
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
