package medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func isValidBST(root *TreeNode) bool {
	var arr []int
	inorderDFS(root, &arr)

	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}

	return true
}

func inorderDFS(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	inorderDFS(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorderDFS(root.Right, arr)
}
