package medium

import (
	"container/heap"
	"math"
)

// (time, vertex)
type tuple [2]int
type minHeap []tuple

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i][0] < m[j][0] }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeap) Push(x any)        { *m = append(*m, x.(tuple)) }
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

// Time: O((v + e)*log(v))
// Space: O(v)
func networkDelayTime(times [][]int, n int, k int) int {
	graph := make(map[int][]tuple, n)
	for _, item := range times {
		u, v, w := item[0], item[1], item[2]
		graph[u] = append(graph[u], tuple{w, v})
	}

	distances := make(map[int]int, n)
	var h minHeap
	heap.Push(&h, tuple{0, k})
	distances[k] = 0

	for h.Len() > 0 {
		t := heap.Pop(&h).(tuple)
		currDst, vertex := t[0], t[1]
		if currDst > distances[vertex] {
			continue
		}

		for _, ngh := range graph[vertex] {
			nghDst, nghVertex := ngh[0], ngh[1]
			newDst := currDst + nghDst
			if newDst < getDst(distances, nghVertex) {
				distances[nghVertex] = newDst
				heap.Push(&h, tuple{newDst, nghVertex})
			}
		}
	}

	if len(distances) != n {
		return -1
	}

	maxTime := math.MinInt
	for _, val := range distances {
		maxTime = max(maxTime, val)
	}

	return maxTime
}

func getDst(distances map[int]int, vertex int) int {
	if dst, exists := distances[vertex]; exists {
		return dst
	}

	return math.MaxInt
}
