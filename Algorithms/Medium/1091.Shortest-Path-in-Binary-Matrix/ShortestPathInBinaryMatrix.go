package medium

type tuple [3]int

// Time: O(n^2)
// Space: O(n)
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 {
		return -1
	}

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	isValid := func(r, c int) bool {
		return r >= 0 && r < n && c >= 0 && c < n && grid[r][c] == 0
	}

	queue := []tuple{{0, 0, 1}}
	grid[0][0] = 1

	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]

		r, c, step := t[0], t[1], t[2]
		if r == n-1 && c == n-1 {
			return step
		}

		for _, d := range directions {
			dr, dc := d[0], d[1]
			newR, newC := r+dr, c+dc
			if isValid(newR, newC) {
				queue = append(queue, tuple{newR, newC, step + 1})
				grid[newR][newC] = 1
			}
		}
	}

	return -1
}
