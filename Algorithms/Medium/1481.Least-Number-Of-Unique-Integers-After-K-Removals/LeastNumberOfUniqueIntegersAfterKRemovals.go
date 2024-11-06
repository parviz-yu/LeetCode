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

// Time: O(nlogn)
// Space: O(n)
func findLeastNumOfUniqueInts(arr []int, k int) int {
	counter := make(map[int]int)
	for _, num := range arr {
		counter[num]++
	}

	var h minHeap
	for _, freq := range counter {
		h = append(h, freq)
	}

	heap.Init(&h)

	var removed_elems int
	for h.Len() > 0 {
		removed_elems += heap.Pop(&h).(int)
		if removed_elems > k {
			return h.Len() + 1
		}
	}

	return 0
}
