package medium

// Time: O(v + e)
// Space: O(v)
func reachableNodes(n int, edges [][]int, restricted []int) int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make(map[int]bool, n)
	for _, r := range restricted {
		visited[r] = true
	}

	visited[0] = true
	queue := []int{0}

	var number int
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		number++
		for _, n := range graph[v] {
			if !visited[n] {
				queue = append(queue, n)
				visited[n] = true
			}
		}
	}

	return number
}
