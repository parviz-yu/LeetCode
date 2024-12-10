package medium

type tuple [3]int

// Time: O(v+e)
// Space: O(v)
func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	isValid := func(r, c int) bool {
		return r >= 0 && r < m && c >= 0 && c < n && maze[r][c] != '+'
	}
	isExit := func(r, c int) bool {
		if r == entrance[0] && c == entrance[1] {
			return false
		}

		return r == 0 || c == 0 || r == m-1 || c == n-1
	}

	row, col := entrance[0], entrance[1]
	queue := []tuple{{row, col, 0}}
	maze[row][col] = '+'

	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		row, col, step := t[0], t[1], t[2]
		if isExit(row, col) {
			return step
		}

		for _, d := range directions {
			dr, dc := d[0], d[1]
			if isValid(row+dr, col+dc) {
				queue = append(queue, tuple{row + dr, col + dc, step + 1})
				maze[row+dr][col+dc] = '+'
			}
		}
	}

	return -1
}
