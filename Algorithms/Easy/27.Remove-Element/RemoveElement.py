# Time Complexity: O(n)
# Space complexity: O(1)
class Solution:
    def removeElement(self, nums: list[int], val: int) -> int:
        l, r = 0, len(nums) - 1
        n = len(nums)

        while l <= r:
            if nums[l] == val:
                if nums[r] != val:
                    nums[l], nums[r] = nums[r], nums[l]
                    l += 1
                r -= 1
                n -= 1
            else:
                l += 1

        return n
