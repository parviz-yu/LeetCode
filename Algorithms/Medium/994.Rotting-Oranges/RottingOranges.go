package medium

type tuple [3]int

// Time: O(m * n)
// Space: O(n)
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	directions := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	isValid := func(r, c int) bool {
		return r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 1
	}

	queue := make([]tuple, 0)
	var freshCount int
	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				freshCount++
			}
			if grid[i][j] == 2 {
				queue = append(queue, tuple{i, j, 0})
			}
		}
	}

	var minutes int
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]

		r, c, currMinute := t[0], t[1], t[2]
		minutes = max(minutes, currMinute)

		for _, d := range directions {
			dr, dc := d[0], d[1]
			newR, newC := r+dr, c+dc
			if isValid(newR, newC) {
				queue = append(queue, tuple{newR, newC, currMinute + 1})
				grid[newR][newC] = 2
				freshCount--
			}
		}

	}

	if freshCount == 0 {
		return minutes
	}

	return -1
}
