package medium

import "sort"

// Time: O(e + v^2*log(v)) = O(v^2*log(v))
// Space: O(v)
func getAncestors(n int, edges [][]int) [][]int {
	graph := make(map[int][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[b] = append(graph[b], a)
	}

	ancestors := make([][]int, n)
	for i := range n {
		stack := []int{i}
		visited := map[int]bool{i: true}
		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, ngh := range graph[node] {
				if !visited[ngh] {
					stack = append(stack, ngh)
					ancestors[i] = append(ancestors[i], ngh)
					visited[ngh] = true
				}
			}
		}

		sort.Ints(ancestors[i])
	}

	return ancestors
}
