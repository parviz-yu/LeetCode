package hard

import (
	"container/heap"
	"sort"
)

type tuple struct {
	capital int
	profit  int
}

type maxHeap []tuple

func (m maxHeap) Len() int      { return len(m) }
func (m maxHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m *maxHeap) Push(x any)   { *m = append(*m, x.(tuple)) }
func (m maxHeap) Less(i, j int) bool {
	if m[i].profit != m[j].profit {
		return m[i].profit > m[j].profit
	}

	return m[i].capital < m[j].capital
}
func (m *maxHeap) Pop() any {
	x := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return x
}

// Time: O(nlog(n))
// Space: O(n)
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	var arr []tuple
	for i := range profits {
		arr = append(arr, tuple{capital: capital[i], profit: profits[i]})
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].capital != arr[j].capital {
			return arr[i].capital < arr[j].capital
		}

		return arr[i].profit > arr[j].profit
	})

	var i int
	var h maxHeap
	for ; k > 0; k-- {

		for i < len(arr) && arr[i].capital <= w {
			heap.Push(&h, arr[i])
			i++
		}

		if h.Len() > 0 {
			t := heap.Pop(&h).(tuple)
			w += t.profit
		}
	}

	return w
}
