package medium

import "container/heap"

// Time: O(nlogk)
// Space: O(n)
func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, n := range nums {
		counter[n]++
	}

	h := &minHeap{}
	for item, count := range counter {
		t := tuple{Item: item, Count: count}
		heap.Push(h, t)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var result []int
	for _, t := range *h {
		result = append(result, t.Item)
	}

	return result
}

type tuple struct {
	Item  int
	Count int
}

type minHeap []tuple

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i].Count < m[j].Count }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeap) Push(x any) {
	*m = append(*m, x.(tuple))
}
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return x
}
