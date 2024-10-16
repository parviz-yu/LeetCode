class Solution:
    # Time: O(n * m)
    # Space: O(1)
    def spiralOrder(self, matrix: list[list[int]]) -> list[int]:
        res = []
        m, n = len(matrix), len(matrix[0])

        dr, dc, r, c = 0, 1, 0, 0

        for _ in range(n * m):
            res.append(matrix[r][c])
            matrix[r][c] = -101

            if matrix[(dr + r) % m][(dc + c) %n] == -101:
                dr, dc = dc, -dr

            r += dr
            c += dc

        return res
