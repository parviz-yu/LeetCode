package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tuple struct {
	Node  *TreeNode
	Depth int
}

// Time: O(n)
// Space: O(n)
func maxDepth_recursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth_recursive(root.Left)
	r := maxDepth_recursive(root.Right)
	return max(l, r) + 1
}

// Time: O(n)
// Space: O(n)
func maxDepth_iterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var result int
	stack := []Tuple{{root, 1}}

	pop := func() (*TreeNode, int) {
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return pair.Node, pair.Depth
	}

	for len(stack) > 0 {
		node, depth := pop()

		if node.Right != nil {
			stack = append(stack, Tuple{node.Right, depth + 1})
		}
		if node.Left != nil {
			stack = append(stack, Tuple{node.Left, depth + 1})
		}

		result = max(result, depth)
	}

	return result
}
