package medium

import "container/heap"

type maxHeap []int

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m maxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *maxHeap) Push(x any)        { *m = append(*m, x.(int)) }
func (m *maxHeap) Pop() any {
	x := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return x
}

// Time: O(klogn)
// Space: O(n)
func minStoneSum(piles []int, k int) int {
	h := maxHeap(piles)
	heap.Init(&h)
	for range k {
		pile := heap.Pop(&h).(int)
		heap.Push(&h, pile-pile/2)
	}

	var total int
	for _, pile := range h {
		total += pile
	}

	return total
}
