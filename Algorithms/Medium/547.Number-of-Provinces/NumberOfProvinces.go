package medium

// Time: O(v^2)
// Space: O(v)
func findCircleNum(isConnected [][]int) int {
	graph := make(map[int][]int)

	n := len(isConnected)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}

	var number int
	visited := make([]bool, n)
	for c := range n {
		if !visited[c] {
			number++

			stack := []int{c}
			visited[c] = true
			for len(stack) > 0 {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				for _, nb := range graph[v] {
					if !visited[nb] {
						stack = append(stack, nb)
						visited[nb] = true
					}
				}
			}
		}
	}

	return number
}
