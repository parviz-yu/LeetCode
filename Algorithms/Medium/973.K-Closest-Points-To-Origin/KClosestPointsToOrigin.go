package medium

import "container/heap"

// Time: O(nlogk)
// Space: O(k)
func kClosest(points [][]int, k int) [][]int {
	m := maxHeap{}
	for _, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		heap.Push(&m, tuple{distance: distance, point: point})
		if m.Len() > k {
			heap.Pop(&m)
		}
	}

	var result [][]int
	for range m.Len() {
		t := heap.Pop(&m).(tuple)
		result = append(result, t.point)
	}

	return result
}

type tuple struct {
	distance int
	point    []int
}

type maxHeap []tuple

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Less(i, j int) bool { return m[i].distance > m[j].distance }
func (m maxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *maxHeap) Push(x any)        { *m = append(*m, x.(tuple)) }
func (m *maxHeap) Pop() any {
	x := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return x
}
