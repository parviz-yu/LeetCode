# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        first, second = 0, 0
        result = []
        while first < len(word1) and second < len(word2):
            if first == second:
                result.append(word1[first])
                first += 1
            else:
                result.append(word2[second])
                second += 1

        if first < len(word1):
            result.extend(word1[first:])

        if second < len(word2):
            result.extend(word2[second:])

        return "".join(result)
