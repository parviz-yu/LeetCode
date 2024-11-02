package medium

import "container/heap"

// Time: O(nlogn)
// Space: O(n)
func frequencySort(s string) string {
	counter := make(map[byte]int)
	for i := range s {
		counter[s[i]]++
	}

	var h maxHeap
	for char, freq := range counter {
		heap.Push(&h, tuple{char: char, freq: freq})
	}

	var result []byte
	for h.Len() > 0 {
		t := heap.Pop(&h).(tuple)
		for range t.freq {
			result = append(result, t.char)
		}
	}

	return string(result)
}

type tuple struct {
	char byte
	freq int
}

type maxHeap []tuple

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Less(i, j int) bool { return m[i].freq > m[j].freq }
func (m maxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(tuple))
}
func (m *maxHeap) Pop() any {
	x := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return x
}
