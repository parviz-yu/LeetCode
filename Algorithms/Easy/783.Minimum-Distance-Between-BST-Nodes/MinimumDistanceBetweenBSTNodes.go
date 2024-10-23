package easy

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func minDiffInBST(root *TreeNode) int {
	var arr []int
	inorder(root, &arr)

	result := math.MaxInt
	for i := 1; i < len(arr); i++ {
		result = min(result, arr[i]-arr[i-1])
	}

	return result
}

func inorder(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}

	inorder(node.Left, arr)
	*arr = append(*arr, node.Val)
	inorder(node.Right, arr)
	return
}
