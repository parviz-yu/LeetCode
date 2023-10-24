# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        s = s.strip().split(" ")
        return len(s[-1])
