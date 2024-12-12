package medium

import (
	"container/heap"
	"math"
)

type maxHeap []int

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m maxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *maxHeap) Push(x any)        { *m = append(*m, x.(int)) }
func (m *maxHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

// Time: O(k*log(n))
// Space: O(n)
func pickGifts(gifts []int, k int) int64 {
	h := maxHeap(gifts)
	heap.Init(&h)

	for range k {
		num := heap.Pop(&h).(int)
		heap.Push(&h, int(math.Sqrt(float64(num))))
	}

	var count int64
	for _, gift := range gifts {
		count += int64(gift)
	}

	return count
}
