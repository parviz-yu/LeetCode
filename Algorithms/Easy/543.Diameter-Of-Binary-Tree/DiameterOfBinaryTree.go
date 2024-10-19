package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func diameterOfBinaryTree(root *TreeNode) int {
	var diameter int
	binTreeDepth(root, &diameter)
	return diameter
}

func binTreeDepth(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	l := binTreeDepth(root.Left, diameter)
	r := binTreeDepth(root.Right, diameter)

	*diameter = max(*diameter, l+r)
	return max(l, r) + 1
}
