package medium

// Time: O(n*(v+e))
// Space: O(n)
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	graph := make(map[int][]int)
	for i := range n - 1 {
		graph[i] = append(graph[i], i+1)
	}

	var result []int
	for _, q := range queries {
		u, v := q[0], q[1]
		graph[u] = append(graph[u], v)

		seen := make(map[int]bool)
		seen[0] = true
		queue := [][]int{{0, 0}}

		for len(queue) > 0 {
			vertex, dst := queue[0][0], queue[0][1]
			queue = queue[1:]
			if vertex == n-1 {
				result = append(result, dst)
				break
			}

			for _, n := range graph[vertex] {
				if !seen[n] {
					queue = append(queue, []int{n, dst + 1})
					seen[n] = true
				}
			}
		}
	}

	return result
}
