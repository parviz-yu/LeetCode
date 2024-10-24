package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}

	var isLeftEq bool
	if isLeftEq = flipEquiv(root1.Left, root2.Left); !isLeftEq {
		root2.Left, root2.Right = root2.Right, root2.Left
	}

	if isLeftEq = flipEquiv(root1.Left, root2.Left); !isLeftEq {
		return false
	}

	var isRightEq bool
	if isRightEq = flipEquiv(root1.Right, root2.Right); !isRightEq {
		root2.Left, root2.Right = root2.Right, root2.Left
	}

	if isRightEq = flipEquiv(root1.Right, root2.Right); !isRightEq {
		return false
	}

	return isLeftEq && isRightEq
}
