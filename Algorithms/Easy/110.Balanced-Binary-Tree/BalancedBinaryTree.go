package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func isBalanced(root *TreeNode) bool {
	result, _ := dfs(root)
	return result
}

func dfs(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}

	leftB, leftH := dfs(node.Left)
	rightB, rightH := dfs(node.Right)

	balanced := leftB && rightB && abs(leftH-rightH) <= 1
	return balanced, max(leftH, rightH) + 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
