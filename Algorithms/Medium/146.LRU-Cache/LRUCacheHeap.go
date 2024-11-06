package medium

import (
	"container/heap"
	"time"
)

type tuple struct {
	timestamp int64
	key       int
	value     int
	index     int // for update
}

type minHeap []*tuple

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i].timestamp < m[j].timestamp }
func (m *minHeap) Push(x any)        { *m = append(*m, x.(*tuple)) }
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return x
}
func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
	m[i].index = i
	m[j].index = j
}

type LRUCacheHeap struct {
	data map[int]*tuple
	h    minHeap
	cap  int
}

func New(capacity int) LRUCacheHeap {
	data := make(map[int]*tuple, capacity)
	h := minHeap{}

	return LRUCacheHeap{
		data: data,
		h:    h,
		cap:  capacity,
	}
}

// Time: O(logn)
func (l *LRUCacheHeap) update(t *tuple) {
	t.timestamp = time.Now().UnixNano()
	heap.Fix(&l.h, t.index)
}

// Time: O(logn)
func (l *LRUCacheHeap) Get(key int) int {
	t, found := l.data[key]
	if !found {
		return -1
	}

	l.update(t)
	return t.value
}

// Time: O(logn)
func (l *LRUCacheHeap) Put(key int, value int) {
	t, found := l.data[key]
	if found {
		t.value = value
		l.update(t)
		return
	}

	if len(l.data) == l.cap {
		t = heap.Pop(&l.h).(*tuple)
		delete(l.data, t.key)
	}

	t = &tuple{
		timestamp: time.Now().UnixNano(),
		key:       key,
		value:     value,
	}

	if l.h.Len() > 0 {
		t.index = l.h.Len()
	}

	heap.Push(&l.h, t)
	l.data[key] = t
}
