package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n)
// Space: O(n)
func isSameTree_recursive(p *TreeNode, q *TreeNode) bool {
	return twoTreesDFS(p, q)
}

func twoTreesDFS(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	left := twoTreesDFS(p.Left, q.Left)
	right := twoTreesDFS(p.Right, q.Right)
	return left && right
}

// Time: O(n)
// Space: O(n)
func isSameTree(p *TreeNode, q *TreeNode) bool {
	stack := []*TreeNode{p, q}

	pop := func() *TreeNode {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return node
	}

	for len(stack) > 0 {
		q, p = pop(), pop()

		if p == nil && q == nil {
			continue
		}

		if p == nil || q == nil || p.Val != q.Val {
			return false
		}

		stack = append(stack, p.Right)
		stack = append(stack, q.Right)
		stack = append(stack, p.Left)
		stack = append(stack, q.Left)
	}

	return true
}
