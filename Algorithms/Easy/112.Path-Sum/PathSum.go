package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tuple struct {
	Node    *TreeNode
	CurrSum int
}

// Time: O(n)
// Space: O(n)
func hasPathSum_recursive(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	leftChild := hasPathSum_recursive(root.Left, targetSum-root.Val)
	rightChild := hasPathSum_recursive(root.Right, targetSum-root.Val)
	return leftChild || rightChild
}

// Time: O(n)
// Space: O(n)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := []Tuple{{root, targetSum}}
	pop := func() (*TreeNode, int) {
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return pair.Node, pair.CurrSum
	}

	for len(stack) > 0 {
		node, currSum := pop()
		if currSum == node.Val && node.Left == nil && node.Right == nil {
			return true
		}

		if node.Left != nil {
			stack = append(stack, Tuple{node.Left, currSum - node.Val})
		}
		if node.Right != nil {
			stack = append(stack, Tuple{node.Right, currSum - node.Val})
		}
	}

	return false
}
