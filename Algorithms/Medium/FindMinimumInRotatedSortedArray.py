# Time complexity: O(logn)
# Space complexity: O(1)
class Solution:
    def findMin(self, nums: list[int]) -> int:
        l, r = 0, len(nums) - 1
        while l < r:
            mid = (l + r) // 2
            if nums[mid] > nums[l]:
                l = mid + 1
            else:
                r = mid

        return nums[l]