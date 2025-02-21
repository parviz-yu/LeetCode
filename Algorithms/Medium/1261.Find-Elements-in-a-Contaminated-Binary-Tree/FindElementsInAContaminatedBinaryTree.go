package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	elems map[int]struct{}
}

// Time: O(n)
// Space: O(n)
func Constructor(root *TreeNode) FindElements {
	f := FindElements{
		elems: make(map[int]struct{}),
	}

	var dfs func(root *TreeNode, val int)
	dfs = func(root *TreeNode, val int) {
		if root == nil {
			return
		}

		f.elems[val] = struct{}{}
		dfs(root.Left, val*2+1)
		dfs(root.Right, val*2+2)
	}

	dfs(root, 0)
	return f
}

// Time: O(1)
func (this *FindElements) Find(target int) bool {
	_, ok := this.elems[target]
	return ok
}
