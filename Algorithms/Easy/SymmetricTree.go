package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time: O(n)
// Space: O(n)
func isSymmetric_recursive(root *helpers.TreeNode) bool {
	return isSymmetricDFS(root.Left, root.Right)
}

func isSymmetricDFS(p, q *helpers.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	leftChild := isSymmetricDFS(p.Left, q.Right)
	rightChild := isSymmetricDFS(p.Right, q.Left)
	return leftChild && rightChild
}

// Time: O(n)
// Space: O(n)
func isSymmetric(root *helpers.TreeNode) bool {
	stack := []*helpers.TreeNode{root.Left, root.Right}

	pop := func() *helpers.TreeNode {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return node
	}

	for len(stack) > 0 {
		right, left := pop(), pop()

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		stack = append(stack, left.Left)
		stack = append(stack, right.Right)
		stack = append(stack, left.Right)
		stack = append(stack, right.Left)
	}

	return true
}
