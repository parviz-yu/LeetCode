package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func mergeTrees(r1 *TreeNode, r2 *TreeNode) *TreeNode {
	if r1 == nil && r2 == nil {
		return nil
	}

	if r1 == nil {
		return r2
	}

	if r2 == nil {
		return r1
	}

	r1.Val += r2.Val

	r1.Left = mergeTrees(r1.Left, r2.Left)
	r1.Right = mergeTrees(r1.Right, r2.Right)
	return r1
}
