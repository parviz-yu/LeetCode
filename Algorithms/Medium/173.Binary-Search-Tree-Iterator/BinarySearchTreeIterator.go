package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	array   []int
	pointer int
}

// Time: O(n)
// Space: O(n)
func Constructor(root *TreeNode) BSTIterator {
	var array []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		array = append(array, root.Val)
		root = root.Right
	}

	return BSTIterator{
		array: array,
	}
}

// Time: O(1)
func (this *BSTIterator) Next() int {
	val := this.array[this.pointer]
	this.pointer++
	return val
}

// Time: O(1)
func (this *BSTIterator) HasNext() bool {
	return this.pointer < len(this.array)
}
