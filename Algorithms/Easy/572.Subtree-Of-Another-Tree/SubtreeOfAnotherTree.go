package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 == nil || r2 == nil || r1.Val != r2.Val {
		return false
	}

	leftChild := isSameTree(r1.Left, r2.Left)
	rightChild := isSameTree(r1.Right, r2.Right)
	return leftChild && rightChild
}
