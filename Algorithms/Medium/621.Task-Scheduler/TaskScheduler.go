package main

import "container/heap"

type tuple struct {
	task int
	time int
}

type maxHeap []int

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m maxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *maxHeap) Push(x any)        { *m = append(*m, x.(int)) }
func (m *maxHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

// Time: O(nlog(n))
// Space: O(1)
func leastInterval(tasks []byte, n int) int {
	counter := make(map[byte]int)
	for _, task := range tasks {
		counter[task]++
	}

	values := make([]int, 0, len(counter))
	for _, val := range counter {
		values = append(values, val)
	}

	h := maxHeap(values)
	heap.Init(&h)

	var queue []tuple
	var time int
	for ; h.Len() > 0 || len(queue) > 0; time++ {
		if h.Len() > 0 {
			cnt := heap.Pop(&h).(int) - 1
			if cnt > 0 {
				queue = append(queue, tuple{cnt, time + n})
			}
		}
		if len(queue) > 0 && queue[0].time == time {
			heap.Push(&h, queue[0].task)
			queue = queue[1:]
		}
	}

	return time
}
