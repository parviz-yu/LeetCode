package medium

// Time: O(v + e)
// Space: O(v)
func canVisitAllRooms(rooms [][]int) bool {
	graph := make(map[int][]int)

	for i, room := range rooms {
		graph[i] = append(graph[i], room...)
	}

	seen := make(map[int]bool, len(graph))
	stack := []int{0}
	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if seen[vertex] {
			continue
		}

		seen[vertex] = true
		stack = append(stack, graph[vertex]...)
	}

	return len(seen) == len(rooms)
}
