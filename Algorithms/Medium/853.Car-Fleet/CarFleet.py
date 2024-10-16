# Time complexity: O(nlogn)
# Space complexity: O(n)
class Solution:
    def carFleet(self, target: int, position: list[int], speed: list[int]) -> int:
        pairs = [[p, s] for p, s in zip(position, speed)]
        stack = []

        for p, s in sorted(pairs)[::-1]:
            stack.append((target - p) / s)
            if len(stack) >= 2 and stack[-1] <= stack[-2]:
                stack.pop()

        return len(stack)