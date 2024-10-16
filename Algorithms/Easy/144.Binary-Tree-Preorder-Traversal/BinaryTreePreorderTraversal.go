package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func preorderTraversal_recursive(root *TreeNode) []int {
	return preorderDFS(root, []int{})
}

func preorderDFS(node *TreeNode, elems []int) []int {
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
func preorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode

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
