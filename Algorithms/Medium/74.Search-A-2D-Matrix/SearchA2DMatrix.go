package medium

// Time Complexity: O(log(m*n))
// Space Complexity: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	up, down := 0, m-1
	for up < down {
		m := (up + down + 1) / 2
		if matrix[m][0] <= target {
			up = m
		} else {
			down = m - 1
		}
	}

	l, r := 0, n-1
	for l <= r {
		m := (l + r) / 2
		if matrix[up][m] == target {
			return true
		}
		if target > matrix[up][m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return false
}
