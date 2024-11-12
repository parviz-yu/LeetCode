package medium

import "container/heap"

type minHeap []string

func (m minHeap) Len() int      { return len(m) }
func (m minHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m minHeap) Less(i, j int) bool {
	if len(m[i]) != len(m[j]) {
		return len(m[i]) < len(m[j])
	}
	return m[i] < m[j]
}
func (m *minHeap) Push(x any) { *m = append(*m, x.(string)) }
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return x
}

// Time: O(nlog(k))
// Space: O(n)
func kthLargestNumber(nums []string, k int) string {
	var h minHeap
	for _, num := range nums {
		heap.Push(&h, num)
		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	return h[0]
}
