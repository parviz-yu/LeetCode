package medium

import "container/heap"

// Time: O(nlogn)
// Space: O(k)
func findKthLargest(nums []int, k int) int {
	var arr Heap
	for _, n := range nums {
		heap.Push(&arr, n)
		if arr.Len() > k {
			heap.Pop(&arr)
		}
	}

	return arr[0]
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	x := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return x
}
