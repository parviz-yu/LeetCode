package medium

import "container/heap"

type minHeap []int

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeap) Push(x any)        { *m = append(*m, x.(int)) }
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

// Time: O(nlog(n))
// Space: O(n)
func minOperations(nums []int, k int) int {
	h := make(minHeap, 0, len(nums))
	for _, num := range nums {
		h = append(h, num)
	}

	heap.Init(&h)

	var count int
	for h.Len() > 1 {
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)

		if x >= k {
			return count
		}

		count++
		heap.Push(&h, min(x, y)*2+max(x, y))
	}

	return count
}
