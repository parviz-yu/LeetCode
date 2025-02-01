package medium

import (
	"container/heap"
	"math"
)

type tuple struct {
	vertex int
	prob   float64
}

type maxHeap []tuple

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Less(i, j int) bool { return m[i].prob > m[j].prob }
func (m maxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *maxHeap) Push(x any)        { *m = append(*m, x.(tuple)) }
func (m *maxHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

// Time: O((v + e)*log(v))
// Space: O(v)
func maxProbability(n int, edges [][]int, succProb []float64, startNode int, endNode int) float64 {
	graph := make(map[int][]tuple, n)
	for i := range edges {
		a, b := edges[i][0], edges[i][1]
		graph[a] = append(graph[a], tuple{vertex: b, prob: succProb[i]})
		graph[b] = append(graph[b], tuple{vertex: a, prob: succProb[i]})
	}

	probs := map[int]float64{startNode: 1}
	var h maxHeap
	heap.Push(&h, tuple{vertex: startNode, prob: 1})

	for h.Len() > 0 {
		t := heap.Pop(&h).(tuple)
		if t.prob < probs[t.vertex] {
			continue
		}

		if t.vertex == endNode {
			return t.prob
		}

		for _, ngh := range graph[t.vertex] {
			newProb := ngh.prob * t.prob
			if newProb > getProb(probs, ngh.vertex) {
				probs[ngh.vertex] = newProb
				heap.Push(&h, tuple{vertex: ngh.vertex, prob: newProb})
			}
		}
	}

	return 0
}

func getProb(probs map[int]float64, vertex int) float64 {
	if prob, exists := probs[vertex]; exists {
		return prob
	}

	return math.SmallestNonzeroFloat64
}
