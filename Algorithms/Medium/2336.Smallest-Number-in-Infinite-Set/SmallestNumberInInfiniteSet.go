package medium

import "container/heap"

type minHeap []int

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeap) Push(x any)        { *m = append(*m, x.(int)) }
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

type SmallestInfiniteSet struct {
	h    minHeap
	seen map[int]bool
}

// Time: O(n)
func Constructor() SmallestInfiniteSet {
	nums := make([]int, 1000)
	seen := make(map[int]bool, 1000)
	for i := range nums {
		nums[i] = i + 1
		seen[i+1] = true
	}

	h := minHeap(nums)
	heap.Init(&h)

	return SmallestInfiniteSet{
		h:    h,
		seen: seen,
	}
}

// Time: O(log(n))
func (s *SmallestInfiniteSet) PopSmallest() int {
	num := heap.Pop(&s.h).(int)
	s.seen[num] = false
	return num
}

// Time: O(log(n))
func (s *SmallestInfiniteSet) AddBack(num int) {
	if !s.seen[num] {
		heap.Push(&s.h, num)
		s.seen[num] = true
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
