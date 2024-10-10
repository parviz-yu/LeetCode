package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

// Time: O(n)
// Space: O(n)
func maxDepth_recursive(root *helpers.TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth_recursive(root.Left)
	r := maxDepth_recursive(root.Right)
	return max(l, r) + 1
}

// Time: O(n)
// Space: O(n)
func maxDepth_iterative(root *helpers.TreeNode) int {
	if root == nil {
		return 0
	}

	type tuple struct {
		node  *helpers.TreeNode
		depth int
	}

	var result int
	stack := []tuple{tuple{root, 1}}

	pop := func() (*helpers.TreeNode, int) {
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return pair.node, pair.depth
	}

	for len(stack) > 0 {
		node, depth := pop()

		if node.Right != nil {
			stack = append(stack, tuple{node.Right, depth + 1})
		}
		if node.Left != nil {
			stack = append(stack, tuple{node.Left, depth + 1})
		}

		result = max(result, depth)
	}

	return result
}
