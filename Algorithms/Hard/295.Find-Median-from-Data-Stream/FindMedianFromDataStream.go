package hard

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

type MedianFinder struct {
	mn minHeap
	mx minHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		mn: minHeap{},
		mx: minHeap{},
	}
}

// Time: O(log n)
// Space: O(n)
func (m *MedianFinder) AddNum(num int) {
	heap.Push(&m.mx, -num)
	heap.Push(&m.mn, -heap.Pop(&m.mx).(int))
	if m.mn.Len() > m.mx.Len() {
		heap.Push(&m.mx, -heap.Pop(&m.mn).(int))
	}
}

// Time: O(1)
// Space: O(n)
func (m *MedianFinder) FindMedian() float64 {
	if m.mn.Len() != m.mx.Len() {
		return float64(-m.mx[0])
	}

	return (float64(m.mn[0]) + float64(-m.mx[0])) / 2
}
