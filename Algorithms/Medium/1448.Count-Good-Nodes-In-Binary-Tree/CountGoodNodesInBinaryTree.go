package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func goodNodes(root *TreeNode) int {
	return dfsHelper(root.Left, root.Val) + dfsHelper(root.Right, root.Val) + 1
}

func dfsHelper(node *TreeNode, currMax int) int {
	if node == nil {
		return 0
	}

	var count int
	if node.Val >= currMax {
		currMax = node.Val
		count++
	}

	count += dfsHelper(node.Left, currMax)
	count += dfsHelper(node.Right, currMax)
	return count
}
