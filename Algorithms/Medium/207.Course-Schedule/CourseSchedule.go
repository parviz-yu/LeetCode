package medium

// Time: O(v + e)
// Space: O(v)
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int, numCourses)
	inDegree := make(map[int]int, numCourses)

	for _, p := range prerequisites {
		a, b := p[0], p[1]
		graph[b] = append(graph[b], a)
		inDegree[a]++
	}

	queue := make([]int, 0)
	for course := range numCourses {
		if inDegree[course] == 0 {
			queue = append(queue, course)
		}
	}

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		numCourses--
		for _, next := range graph[course] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return numCourses == 0
}
