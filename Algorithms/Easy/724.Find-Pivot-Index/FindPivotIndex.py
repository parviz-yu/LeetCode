# Time complexity: O(n)
class Solution:
    # Space complexity: O(n)
    def pivotIndex(self, nums: list[int]) -> int:
        total_sum = sum(nums)

        left_sum = 0
        for i in range(len(nums)):
            right_sum = total_sum - left_sum - nums[i]
            if left_sum == right_sum:
                return i

            left_sum += nums[i]

        return -1

    # Space complexity: O(n)
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
