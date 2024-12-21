package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	dfs(root.Left, root.Right, 1)
	return root
}

func dfs(left, right *TreeNode, depth int) {
	if left == nil && right == nil {
		return
	}

	if depth%2 == 1 {
		left.Val, right.Val = right.Val, left.Val
	}

	dfs(left.Left, right.Right, depth+1)
	dfs(left.Right, right.Left, depth+1)
}
