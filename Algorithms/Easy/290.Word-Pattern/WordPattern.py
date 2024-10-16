# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        split = s.split(" ")
        if len(split) != len(pattern):
            return False

        pToS = {}
        sToP = {}
        for i in range(len(pattern)):
            if (pattern[i] in pToS and pToS[pattern[i]] != split[i]) or (
                split[i] in sToP and sToP[split[i]] != pattern[i]
            ):
                return False
            pToS[pattern[i]] = split[i]
            sToP[split[i]] = pattern[i]

        return True
