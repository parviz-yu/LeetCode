# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def arraySign(self, nums: list[int]) -> int:
        is_positive = True
        for num in nums:
            if num == 0:
                return 0

            if num < 0:
                is_positive = not is_positive

        if is_positive:
            return 1

        return -1
