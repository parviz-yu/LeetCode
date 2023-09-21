# Time Complexity: O(log(m*n))
# Space Complexity: O(1)
class Solution:
    def searchMatrix(self, matrix: list[list[int]], target: int) -> bool:
        m, n = len(matrix), len(matrix[0])
    
        # Finding last smallest element in col #1
        l_col, r_col = 0, m - 1
        while l_col < r_col:
            mid_col = (l_col + r_col + 1) // 2
            if matrix[mid_col][0] <= target:
                l_col = mid_col
            else:
                r_col = mid_col - 1

        l_row, r_row = 0, n - 1
        while l_row <= r_row:
            mid_row = (l_row + r_row) // 2
            if matrix[l_col][mid_row] == target:
                return True
            if matrix[l_col][mid_row] < target:
                l_row = mid_row + 1
            else:
                r_row = mid_row - 1

        return False