package medium

import "container/heap"

type tuple struct {
	idx   int
	value int
}

type minHeap []tuple

func (m minHeap) Len() int      { return len(m) }
func (m minHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m *minHeap) Push(x any)   { *m = append(*m, x.(tuple)) }
func (m minHeap) Less(i, j int) bool {
	if m[i].value == m[j].value {
		return m[i].idx < m[j].idx
	}

	return m[i].value < m[j].value
}
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

// Time: O(n*log(n))
// Space: O(n)
func findScore(nums []int) int64 {
	tuples := make([]tuple, 0, len(nums))
	for i, n := range nums {
		tuples = append(tuples, tuple{idx: i, value: n})
	}
	h := minHeap(tuples)
	heap.Init(&h)

	var score int
	for h.Len() > 0 {
		t := heap.Pop(&h).(tuple)
		if nums[t.idx] > 0 {
			score += nums[t.idx]
			nums[t.idx] = -1
			if t.idx-1 >= 0 && t.idx-1 < len(nums) {
				nums[t.idx-1] = -1
			}
			if t.idx+1 >= 0 && t.idx+1 < len(nums) {
				nums[t.idx+1] = -1
			}
		}
	}

	return int64(score)
}
