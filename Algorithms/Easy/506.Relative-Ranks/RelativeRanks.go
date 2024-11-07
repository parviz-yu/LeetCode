package easy

import (
	"container/heap"
	"strconv"
)

type tuple struct {
	score int
	place int
}

type minHeap []tuple

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i].score < m[j].score }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeap) Push(x any)        { *m = append(*m, x.(tuple)) }
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return x
}

// Time: O(nlogn)
// Space: O(n)
func findRelativeRanks(score []int) []string {
	var h minHeap
	for i, s := range score {
		heap.Push(&h, tuple{score: s, place: i})
	}

	result := make([]string, len(score))
	for k := h.Len(); h.Len() > 0; k-- {
		t := heap.Pop(&h).(tuple)
		if k == 1 {
			result[t.place] = "Gold Medal"
		} else if k == 2 {
			result[t.place] = "Silver Medal"
		} else if k == 3 {
			result[t.place] = "Bronze Medal"
		} else {
			result[t.place] = strconv.Itoa(h.Len() + 1)
		}
	}

	return result
}
