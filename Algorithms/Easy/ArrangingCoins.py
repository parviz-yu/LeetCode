# Time complexity: O(logn)
# Space complexity: O(1)
class Solution:
    def arrangeCoins(self, n: int) -> int:
        l, r = 1, n
        while l < r:
            m = (l + r + 1) // 2
            # Sum of m terms of an Arithmetic Progression
            if (m + 1) * m // 2 <= n:
                l = m
            else:
                r = m - 1

        return l
