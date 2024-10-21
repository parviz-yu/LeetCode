package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(m + n)
// Space: O(m + n)
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1 := dfs(root1)
	arr2 := dfs(root2)

	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return nil
	}

	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}

	l := dfs(node.Left)
	r := dfs(node.Right)
	return append(l, r...)
}
