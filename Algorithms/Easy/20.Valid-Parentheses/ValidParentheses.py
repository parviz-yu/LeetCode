# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def isValid(self, s: str) -> bool:
        stack = []

        for c in s:
            if c in "([{":
                stack.append(c)
            else:
                if len(stack) == 0:
                    return False

                pair = stack.pop() + c
                if pair not in ["()", "[]", "{}"]:
                    return False

        if len(stack) != 0:
            return False

        return True      