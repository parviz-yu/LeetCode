package medium

type tuple struct {
	row, col int
}

// Time: O(m*n)
// Space: O(m*n)
func numIslands(grid [][]byte) int {
	var number int

	m, n := len(grid), len(grid[0])
	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	isValid := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == '1'
	}

	for i := range m {
		for j := range n {
			if grid[i][j] == '1' {
				number++
				stack := []tuple{{i, j}}
				for len(stack) > 0 {
					t := stack[len(stack)-1]
					stack = stack[:len(stack)-1]

					if isValid(t.row, t.col) {
						grid[t.row][t.col] = '0' // mark as visited
						for _, d := range directions {
							stack = append(stack, tuple{t.row + d[0], t.col + d[1]})
						}
					}
				}
			}
		}
	}

	return number
}
