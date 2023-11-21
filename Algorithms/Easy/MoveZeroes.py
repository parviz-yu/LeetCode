# Time complexity: O(n)
class Solution:
    # Space complexity: O(n)
    def moveZeroes(self, nums: list[int]) -> None:
        nums_cp = nums[:]

        i, j = 0, len(nums) - 1
        for num in nums_cp:
            if num != 0:
                nums[i] = num
                i += 1
            else:
                nums[j] = num
                j -= 1

    # Space complexity: O(1)
    def moveZeroes_1(self, nums: list[int]) -> None:
        slow = 0
        for fast in range(len(nums)):
            if nums[fast]:
                nums[slow], nums[fast] = nums[fast], nums[slow]
                slow += 1
