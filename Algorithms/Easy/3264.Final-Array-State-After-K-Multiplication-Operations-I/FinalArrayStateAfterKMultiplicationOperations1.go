package medium

import "container/heap"

type minHeap [][2]int // idx, value

func (m minHeap) Len() int      { return len(m) }
func (m minHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m minHeap) Less(i, j int) bool {
	if m[i][1] == m[j][1] {
		return m[i][0] < m[j][0]
	}

	return m[i][1] < m[j][1]
}
func (m *minHeap) Push(x any) { *m = append(*m, x.([2]int)) }
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

// Time: O(n*log(n))
// Space: O(n)
func getFinalState(nums []int, k int, multiplier int) []int {
	var h minHeap
	for i, num := range nums {
		heap.Push(&h, [2]int{i, num})
	}

	for range k {
		t := heap.Pop(&h).([2]int)
		t[1] *= multiplier
		heap.Push(&h, t)
	}

	result := make([]int, len(nums))
	for h.Len() > 0 {
		t := heap.Pop(&h).([2]int)
		result[t[0]] = t[1]
	}

	return result
}
