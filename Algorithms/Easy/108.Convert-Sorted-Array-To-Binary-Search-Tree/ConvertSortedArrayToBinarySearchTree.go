package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Time: O(n)
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	i := len(nums) / 2
	node := &TreeNode{Val: nums[i]}
	node.Left = sortedArrayToBST(nums[:i])
	node.Right = sortedArrayToBST(nums[i+1:])
	return node
}
