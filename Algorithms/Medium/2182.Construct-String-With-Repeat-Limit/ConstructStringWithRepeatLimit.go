package medium

import "container/heap"

type maxHeap []byte

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m maxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *maxHeap) Push(x any)        { *m = append(*m, x.(byte)) }
func (m *maxHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

// Time: O(n*log(n))
// Space: O(n)
func repeatLimitedString(s string, repeatLimit int) string {
	counter := make(map[byte]int, 26)
	for i := range s {
		char := s[i]
		counter[char]++
	}

	var h maxHeap
	for char := range counter {
		heap.Push(&h, char)
	}

	buf := make([]byte, 0, len(s))
	for h.Len() > 0 {
		char := heap.Pop(&h).(byte)
		if len(buf) > 0 && buf[len(buf)-1] == char {
			break
		}

		k := repeatLimit
		for k > 0 && counter[char] > 0 {
			buf = append(buf, char)
			k--
			counter[char]--
		}

		if counter[char] > 0 && h.Len() > 0 {
			tempChar := heap.Pop(&h).(byte)
			buf = append(buf, tempChar)
			counter[tempChar]--

			heap.Push(&h, char)
			if counter[tempChar] > 0 {
				heap.Push(&h, tempChar)
			}
		}
	}

	return string(buf)
}
