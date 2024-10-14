package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time: O(n)
// Space: O(n)
func searchBST_recursive(root *helpers.TreeNode, val int) *helpers.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if val < root.Val {
		return searchBST_recursive(root.Left, val)
	}

	return searchBST_recursive(root.Right, val)
}

// Time: O(n)
// Space: O(n)
func searchBST(root *helpers.TreeNode, val int) *helpers.TreeNode {
	stack := []*helpers.TreeNode{root}
	pop := func() *helpers.TreeNode {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return node
	}

	for len(stack) > 0 {
		root = pop()

		if root == nil {
			continue
		}

		if root.Val == val {
			return root
		}

		if val < root.Val {
			stack = append(stack, root.Left)
		} else {
			stack = append(stack, root.Right)
		}
	}

	return nil
}
