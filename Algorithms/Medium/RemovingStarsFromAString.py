# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def removeStars(self, s: str) -> str:
        stack = []

        for ch in s:
            if ch == "*":
                stack.pop()
            else:
                stack.append(ch)

        return "".join(stack)
