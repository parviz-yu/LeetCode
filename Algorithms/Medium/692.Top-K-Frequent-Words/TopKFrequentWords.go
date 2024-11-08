package medium

import "container/heap"

type tuple struct {
	word string
	freq int
}

type minHeap []tuple

func (m minHeap) Len() int      { return len(m) }
func (m minHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m minHeap) Less(i, j int) bool {
	if m[i].freq != m[j].freq {
		return m[i].freq < m[j].freq
	}
	return m[i].word > m[j].word
}
func (m *minHeap) Push(x any) {
	*m = append(*m, x.(tuple))
}
func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return x
}

// Time: O(nlogk)
// Space: O(n)
func topKFrequent(words []string, k int) []string {
	counter := make(map[string]int)
	for _, word := range words {
		counter[word]++
	}

	var h minHeap
	for word, freq := range counter {
		heap.Push(&h, tuple{word: word, freq: freq})
		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	freqWords := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		t := heap.Pop(&h).(tuple)
		freqWords[i] = t.word
	}

	return freqWords
}
