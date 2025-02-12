package medium

// Time: O(n^2)
// Space: O(n)
func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	graph := make(map[int][]int, n)
	for i := range n {
		for j := range n {
			if i == j {
				continue
			}

			if canCover(bombs[i], bombs[j]) {
				graph[i] = append(graph[i], j)
			}
		}
	}

	maxCount := 0
	for i := range n {
		if _, exists := graph[i]; !exists {
			continue
		}

		visited := make(map[int]bool, len(graph))
		stack := []int{i}
		visited[i] = true
		count := 1

		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for _, ngh := range graph[node] {
				if !visited[ngh] {
					stack = append(stack, ngh)
					visited[ngh] = true
					count++
				}
			}
		}

		maxCount = max(maxCount, count)
	}

	if maxCount == 0 {
		return 1
	}

	return maxCount
}

func canCover(first, second []int) bool {
	x1, y1, r1 := first[0], first[1], first[2]
	x2, y2 := second[0], second[1]
	distance := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	return distance <= r1*r1
}
