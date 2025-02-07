package medium

// Time: O(v + e)
// Space: O(v)
func findOrder(numCourses int, prerequisites [][]int) []int {
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

	result := make([]int, 0, numCourses)
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		numCourses--
		result = append(result, course)
		for _, next := range graph[course] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	if numCourses != 0 {
		return []int{}
	}

	return result
}
