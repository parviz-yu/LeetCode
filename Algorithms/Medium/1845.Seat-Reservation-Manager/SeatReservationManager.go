package medium

import "container/heap"

type minHeap []int

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeap) Push(x any)        { *m = append(*m, x.(int)) }
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return x
}

type SeatManager struct {
	h minHeap
}

// Time: O(n)
func Constructor(n int) SeatManager {
	h := minHeap(make([]int, 0, n))
	for i := range n {
		h = append(h, i+1)
	}

	heap.Init(&h)
	return SeatManager{h: h}
}

// Time: O(log(n))
func (s *SeatManager) Reserve() int {
	return heap.Pop(&s.h).(int)
}

// Time: O(log(n))
func (s *SeatManager) Unreserve(seatNumber int) {
	heap.Push(&s.h, seatNumber)
}
