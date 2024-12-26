package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func maxLevelSum(root *TreeNode) int {
	maxVal := root.Val
	lvl := 1

	var levelNum int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)

		var levelSum int
		for i := range n {
			root = queue[i]
			levelSum += root.Val
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}

		queue = queue[n:]
		levelNum++
		if levelSum > maxVal {
			maxVal = levelSum
			lvl = levelNum
		}
	}

	return lvl
}
