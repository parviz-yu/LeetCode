package medium

// Time: O(v + e)
// Space: O(v)
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	inDegree := make([]int, n)
	for _, edge := range edges {
		to := edge[1]
		inDegree[to]++
	}

	var vertices []int
	for i, v := range inDegree {
		if v == 0 {
			vertices = append(vertices, i)
		}
	}

	return vertices
}
