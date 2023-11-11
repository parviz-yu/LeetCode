# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def pivotIndex(self, nums: list[int]) -> int:
        prefix = [0] * (len(nums) + 1)
        for i in range(len(nums)):
            prefix[i + 1] = prefix[i] + nums[i]

        l, r = 0, len(nums)
        for m in range(len(nums)):
            left_sum = prefix[m] - prefix[l]
            right_sum = prefix[r] - prefix[m + 1]
            if left_sum == right_sum:
                return m

        return -1
