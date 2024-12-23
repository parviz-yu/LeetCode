package medium

import (
	"container/list"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n*log(n))
// Space: O(n)
func minimumOperations(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)

	var swaps int
	for queue.Len() > 0 {
		var level []int
		for range queue.Len() {
			root = queue.Remove(queue.Front()).(*TreeNode)
			level = append(level, root.Val)
			if root.Left != nil {
				queue.PushBack(root.Left)
			}
			if root.Right != nil {
				queue.PushBack(root.Right)
			}
		}

		swaps += getSwaps(level)
	}

	return swaps
}

func getSwaps(arr []int) int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	slices.SortFunc(sortedArr, func(a, b int) int {
		return a - b
	})

	elemPos := make(map[int]int, len(arr))
	for i, elem := range arr {
		elemPos[elem] = i
	}

	var swaps int
	for i := range arr {
		if arr[i] != sortedArr[i] {
			corrIdx := elemPos[sortedArr[i]]
			elemPos[arr[i]] = corrIdx
			arr[i], arr[corrIdx] = arr[corrIdx], arr[i]
			swaps++
		}
	}

	return swaps
}
