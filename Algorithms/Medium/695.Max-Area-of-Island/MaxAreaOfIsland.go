package medium

type tuple struct {
	row, col int
}

func maxAreaOfIsland(grid [][]int) int {
	var maxArea int

	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	m, n := len(grid), len(grid[0])
	isValid := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 1
	}

	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				var currArea int
				stack := []tuple{{i, j}}
				grid[i][j] = 0 // mark as visited

				for len(stack) > 0 {
					t := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					currArea++

					for _, d := range directions {
						dr := t.row + d[0]
						dc := t.col + d[1]
						if isValid(dr, dc) {
							stack = append(stack, tuple{dr, dc})
							grid[dr][dc] = 0 // mark as visited
						}
					}
				}

				maxArea = max(maxArea, currArea)
			}

		}
	}

	return maxArea
}
