# Time complexity: O(logn)
# Space complexity: O(1)
class Solution:
    def searchInsert(self, nums: list[int], target: int) -> int:
        l, r = 0, len(nums) - 1
        while l < r:
            m = (l + r) // 2
            if target <= nums[m]:
                r = m
            else:
                l = m + 1

        if target > nums[l]:
            l += 1

        return l
