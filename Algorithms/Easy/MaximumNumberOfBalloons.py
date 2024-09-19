# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        pattern = self.counter("balloon")
        occur = self.counter(text)

        res = len(text)
        for k in pattern:
            res = min(res, occur.get(k, 0) // pattern[k])

        return res

    def counter(self, word: str) -> dict[str, int]:
        res = {}
        for c in word:
            res[c] = res.get(c, 0) + 1

        return res
