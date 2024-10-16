# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def longestCommonPrefix(self, strs: list[str]) -> str:
        smallest = min(strs, key=len)
        for i, c in enumerate(smallest):
            for s in strs:
                if s[i] != c:
                    return smallest[:i]

        return smallest
