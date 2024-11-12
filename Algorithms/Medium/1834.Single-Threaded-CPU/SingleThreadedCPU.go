package medium

import (
	"container/heap"
	"sort"
)

type tuple struct {
	index       int
	enqueueTime int
	processTime int
}

type minHeap []tuple

func (m minHeap) Len() int      { return len(m) }
func (m minHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m minHeap) Less(i, j int) bool {
	if m[i].processTime != m[j].processTime {
		return m[i].processTime < m[j].processTime
	}

	return m[i].index < m[j].index
}
func (m *minHeap) Push(x any) { *m = append(*m, x.(tuple)) }
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return x
}

func getOrder(tasks [][]int) []int {
	var sorted []tuple
	for i, task := range tasks {
		sorted = append(sorted, tuple{i, task[0], task[1]})
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].enqueueTime != sorted[j].enqueueTime {
			return sorted[i].enqueueTime < sorted[j].enqueueTime
		}

		if sorted[i].processTime != sorted[j].processTime {
			return sorted[i].processTime < sorted[j].processTime
		}

		return sorted[i].index < sorted[j].index
	})

	var h minHeap
	var result []int
	var ptr int
	minEnqTime := sorted[ptr].enqueueTime
	n := len(sorted)
	for len(result) < n {
		for ptr < n {
			if sorted[ptr].enqueueTime > minEnqTime {
				break
			}

			heap.Push(&h, sorted[ptr])
			ptr++
		}

		if h.Len() > 0 {
			t := heap.Pop(&h).(tuple)
			minEnqTime += t.processTime
			result = append(result, t.index)
		} else if ptr < n {
			minEnqTime = sorted[ptr].enqueueTime
		}
	}

	return result
}
