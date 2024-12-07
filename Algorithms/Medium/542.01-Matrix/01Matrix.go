package medium

// Time: O(v^2)
// Space: O(v)
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	queue := make([][3]int, 0)
	visited := make(map[[2]int]bool)

	isValid := func(r, c int) bool {
		return r >= 0 && r < m && c >= 0 && c < n && !visited[[2]int{r, c}]
	}

	for i := range m {
		for j := range n {
			if mat[i][j] == 0 {
				queue = append(queue, [3]int{i, j, 1})
				visited[[2]int{i, j}] = true
			}
		}
	}

	for len(queue) > 0 {
		tuple := queue[0]
		queue = queue[1:]

		r, c, depth := tuple[0], tuple[1], tuple[2]
		for _, d := range directions {
			dr, dc := d[0], d[1]
			if isValid(r+dr, c+dc) {
				mat[r+dr][c+dc] = depth
				queue = append(queue, [3]int{r + dr, c + dc, depth + 1})
				visited[[2]int{r + dr, c + dc}] = true
			}
		}
	}

	return mat
}
