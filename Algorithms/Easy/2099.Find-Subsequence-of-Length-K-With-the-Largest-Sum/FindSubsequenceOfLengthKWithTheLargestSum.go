package easy

import (
	"container/heap"
	"math"
)

type tuple struct {
	num int
	idx int
}

type minHeap []tuple

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i].num < m[j].num }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeap) Push(x any)        { *m = append(*m, x.(tuple)) }
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return x
}

// Time: O(nlogk)
// Time: O(k)
func maxSubsequence(nums []int, k int) []int {
	var h minHeap
	for idx, num := range nums {
		heap.Push(&h, tuple{num: num, idx: idx})
		if h.Len() > k {
			t := heap.Pop(&h).(tuple)
			nums[t.idx] = math.MinInt
		}
	}

	result := make([]int, 0, k)
	for _, num := range nums {
		if num != math.MinInt {
			result = append(result, num)
		}
	}

	return result
}
