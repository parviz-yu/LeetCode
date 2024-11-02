package medium

import (
	"container/heap"
)

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *Heap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// Space: O(k)
type KthLargest struct {
	MinHeap *Heap
	k       int
}

// Time: O(logk)
func Constructor(k int, nums []int) KthLargest {
	obj := KthLargest{
		MinHeap: &Heap{},
		k:       k,
	}

	for _, item := range nums {
		obj.Add(item)
	}

	return obj
}

// Time: O(logk)
func (this *KthLargest) Add(val int) int {
	heap.Push(this.MinHeap, val)
	if this.MinHeap.Len() > this.k {
		heap.Pop(this.MinHeap)
	}

	return (*this.MinHeap)[0]
}
