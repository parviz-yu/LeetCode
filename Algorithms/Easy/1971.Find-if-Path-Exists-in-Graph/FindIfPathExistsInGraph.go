package medium

// Time: O(v + e)
// Space: O(v)
func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n)
	stack := []int{source}
	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if vertex == destination {
			return true
		}

		for _, c := range graph[vertex] {
			if !visited[c] {
				stack = append(stack, c)
				visited[c] = true
			}
		}
	}

	return false
}

// Time: O(v + e)
// Space: O(v)
func validPath_recursive(n int, edges [][]int, source int, destination int) bool {
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n)
	var dfs func(int) bool
	dfs = func(vertex int) bool {
		if vertex == destination {
			return true
		}

		for _, n := range graph[vertex] {
			if !visited[n] {
				visited[n] = true
				if dfs(n) {
					return true
				}
			}
		}

		return false
	}

	return dfs(source)
}
