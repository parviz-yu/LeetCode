# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        pattern = "balloon"
        occurances = {c: 0 for c in pattern}

        for char in text:
            if char in occurances:
                occurances[char] += 1

        occurances["l"] //= 2
        occurances["o"] //= 2

        return min(occurances.values())
