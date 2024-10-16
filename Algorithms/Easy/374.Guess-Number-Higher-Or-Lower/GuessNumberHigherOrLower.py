# Time complexity: O(logn)
# Space complexity: O(1)
class Solution:
    def guessNumber(self, n: int) -> int:
        l, r = 1, n
        while l <= r:
            m = (l + r) // 2
            ans = guess(m)
            if ans == 0:
                return m
            if ans == -1:
                r = m - 1
            else:
                l = m + 1
