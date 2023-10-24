# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def getConcatenation(self, nums: list[int]) -> list[int]:
        n = len(nums)
        ans = [0] * n * 2
        for i in range(n):
            ans[i] = ans[i + n] = nums[i]

        return ans
