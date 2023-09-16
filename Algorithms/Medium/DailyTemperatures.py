# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def dailyTemperatures(self, temperatures: list[int]) -> list[int]:
        result = [0] * len(temperatures)
        stack = []

        for i, temp in enumerate(temperatures):
            while stack and temperatures[stack[-1]] < temp:
                prev_idx = stack.pop()
                result[prev_idx] = i - prev_idx
                
            stack.append(i)

        return result