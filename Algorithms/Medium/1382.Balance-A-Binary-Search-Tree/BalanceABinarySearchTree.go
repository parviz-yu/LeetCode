package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func balanceBST(root *TreeNode) *TreeNode {
	return buildBST(inorderTraverse(root))
}

func inorderTraverse(root *TreeNode) []int {
	var result []int

	var stack []*TreeNode
	for len(stack) != 0 || root != nil {
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

func buildBST(array []int) *TreeNode {
	if len(array) == 0 {
		return nil
	}

	idx := len(array) / 2
	root := &TreeNode{Val: array[idx]}
	root.Left = buildBST(array[:idx])
	root.Right = buildBST(array[idx+1:])
	return root
}
