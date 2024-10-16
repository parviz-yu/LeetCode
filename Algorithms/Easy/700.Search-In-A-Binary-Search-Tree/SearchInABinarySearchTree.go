package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func searchBST_recursive(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if val < root.Val {
		return searchBST_recursive(root.Left, val)
	}

	return searchBST_recursive(root.Right, val)
}

// Time: O(n)
// Space: O(n)
func searchBST(root *TreeNode, val int) *TreeNode {
	stack := []*TreeNode{root}
	pop := func() *TreeNode {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return node
	}

	for len(stack) > 0 {
		root = pop()

		if root == nil {
			continue
		}

		if root.Val == val {
			return root
		}

		if val < root.Val {
			stack = append(stack, root.Left)
		} else {
			stack = append(stack, root.Right)
		}
	}

	return nil
}
