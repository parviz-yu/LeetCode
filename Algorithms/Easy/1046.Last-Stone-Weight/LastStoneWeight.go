package easy

import "container/heap"

// Time: O(nlogn)
// Space: O(n)
func lastStoneWeight(stones []int) int {
	m := maxHeap(stones)
	heap.Init(&m) // heapify

	for m.Len() > 1 {
		x, y := heap.Pop(&m).(int), heap.Pop(&m).(int)
		if x != y {
			heap.Push(&m, x-y)
		}
	}

	if m.Len() == 0 {
		return 0
	}

	return m[0]
}

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
