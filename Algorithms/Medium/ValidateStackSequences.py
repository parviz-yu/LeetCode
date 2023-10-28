# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def validateStackSequences(self, pushed: list[int], popped: list[int]) -> bool:
        stack = []
        i, j = 0, 0

        while i < len(pushed):
            if stack and stack[-1] == popped[j]:
                stack.pop()
                j += 1
            else:
                stack.append(pushed[i])
                i += 1

        while j < len(popped):
            if stack.pop() != popped[j]:
                return False
            j += 1

        return True
