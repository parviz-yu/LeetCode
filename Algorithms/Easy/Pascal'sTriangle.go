package easy

// Time complexity: O(n^2)
// Space complexity: O(1)
func generate(numRows int) [][]int {
	triangle := make([][]int, 0, numRows)
	triangle = append(triangle, []int{1})

	for i := 0; i < numRows-1; i++ {
		row := make([]int, 0, len(triangle[i])+1)
		temp := 0

		for _, val := range triangle[i] {
			row = append(row, temp+val)
			temp = val
		}

		row = append(row, 1)
		triangle = append(triangle, row)
	}

	return triangle
}
