package medium

import (
	"container/list"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(nlogn)
// Space: O(n)
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	queue := list.New()
	queue.PushBack(root)

	var sums []int
	for queue.Len() > 0 {

		var levelSum int
		for range queue.Len() {
			root = queue.Remove(queue.Front()).(*TreeNode)
			levelSum += root.Val

			if root.Left != nil {
				queue.PushBack(root.Left)
			}
			if root.Right != nil {
				queue.PushBack(root.Right)
			}
		}

		sums = append(sums, levelSum)
	}

	// desc sort
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] >= sums[j]
	})

	if k > len(sums) {
		return -1
	}

	return int64(sums[k-1])
}
