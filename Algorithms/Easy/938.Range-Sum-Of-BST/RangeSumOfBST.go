package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	l := rangeSumBST(root.Left, low, high)
	r := rangeSumBST(root.Right, low, high)

	if isInRange(root.Val, low, high) {
		l += root.Val
	}

	return l + r
}

func isInRange(val, low, high int) bool {
	if val >= low && val <= high {
		return true
	}

	return false
}
