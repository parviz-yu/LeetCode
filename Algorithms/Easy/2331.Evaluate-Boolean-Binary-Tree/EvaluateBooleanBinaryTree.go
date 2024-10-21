package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}

	l := evaluateTree(root.Left)
	r := evaluateTree(root.Right)
	return evaluation(root.Val, l, r)
}

func evaluation(operand int, a, b bool) bool {
	if operand == 2 {
		return a || b
	}

	return a && b
}
