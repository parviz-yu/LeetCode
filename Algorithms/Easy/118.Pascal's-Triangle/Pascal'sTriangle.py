# Time complexity: O(n^2)
# Space complexity: O(1)
class Solution:
    def generate(self, numRows: int) -> list[list[int]]:
        triangle = [[1]]
        for i in range(numRows - 1):
            temp = 0
            row = []
            for j in range(len(triangle[i])):
                row.append(temp + triangle[i][j])
                temp = triangle[i][j]

            row.append(1)
            triangle.append(row)

        return triangle
