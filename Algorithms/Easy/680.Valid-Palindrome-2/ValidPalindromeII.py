# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def validPalindrome(self, s: str) -> bool:
        l, r = 0, len(s) - 1
        while l < r:
            if s[l] != s[r]:
                return self.isValid(s, l + 1, r) or self.isValid(s, l, r - 1)

            l += 1
            r -= 1

        return True

    def isValid(self, s: str, l: int, r: int) -> bool:
        while l < r:
            if s[l] != s[r]:
                return False

            l += 1
            r -= 1

        return True
