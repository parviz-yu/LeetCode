package medium

import (
	"container/heap"
	"fmt"
	"sort"
)

type tuple1 struct {
	fraction float64
	ints     []int
}

// Time: O(n^2)
// Space: O(n)
func kthSmallestPrimeFraction(arr []int, k int) []int {
	var toSort []tuple1
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			t := tuple1{
				fraction: float64(arr[i]) / float64(arr[j]),
				ints:     []int{arr[i], arr[j]},
			}
			toSort = append(toSort, t)
		}
	}

	sort.Slice(toSort, func(i, j int) bool {
		return toSort[i].fraction < toSort[j].fraction
	})

	return toSort[k-1].ints
}

// ----------------

type tuple struct {
	fraction float64
	numIdx   int
	denomIdx int
}

type minHeap []tuple

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i].fraction < m[j].fraction }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeap) Push(x any)        { *m = append(*m, x.(tuple)) }
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return x
}

// Time: O((n+k)logn)
// Space: O(n)
func kthSmallestPrimeFraction_(arr []int, k int) []int {
	var h minHeap
	n := len(arr)
	for i := 0; i < n-1; i++ {
		heap.Push(&h, tuple{
			fraction: float64(arr[i]) / float64(arr[n-1]),
			numIdx:   i,
			denomIdx: n - 1,
		})
	}

	fmt.Println(h)

	for range k - 1 {
		t := heap.Pop(&h).(tuple)
		t.denomIdx--
		if t.numIdx < t.denomIdx {
			heap.Push(&h, tuple{
				fraction: float64(arr[t.numIdx]) / float64(arr[t.denomIdx]),
				numIdx:   t.numIdx,
				denomIdx: t.denomIdx,
			})
		}
	}

	result := heap.Pop(&h).(tuple)
	return []int{arr[result.numIdx], arr[result.denomIdx]}
}
