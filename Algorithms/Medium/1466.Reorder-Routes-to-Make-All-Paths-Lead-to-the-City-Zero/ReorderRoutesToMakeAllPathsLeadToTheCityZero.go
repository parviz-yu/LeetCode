package medium

// Time: O(v + e)
// Space: O(v)
func minReorder(n int, connections [][]int) int {
	hasDirections := make(map[[2]int]bool, len(connections))
	graph := make(map[int][]int, n)
	for _, edge := range connections {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
		hasDirections[[2]int{a, b}] = true
	}

	visited := make([]bool, n)
	visited[0] = true
	stack := []int{0}
	var turns int

	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, n := range graph[v] {
			if !visited[n] {
				if hasDirections[[2]int{v, n}] {
					turns++
				}
				visited[n] = true
				stack = append(stack, n)
			}
		}
	}

	return turns
}
