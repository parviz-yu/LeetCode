package medium

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func tree2str(root *TreeNode) string {
	return strings.Join(dfs(root, []string{}), "")
}

func dfs(node *TreeNode, buf []string) []string {
	if node == nil {
		buf = append(buf, "")
		return buf
	}

	buf = append(buf, strconv.Itoa(node.Val))

	if node.Left != nil || node.Right != nil {
		buf = append(buf, "(")
		buf = dfs(node.Left, buf)
		buf = append(buf, ")")
	}

	if node.Right != nil {
		buf = append(buf, "(")
		buf = dfs(node.Right, buf)
		buf = append(buf, ")")
	}

	return buf
}
