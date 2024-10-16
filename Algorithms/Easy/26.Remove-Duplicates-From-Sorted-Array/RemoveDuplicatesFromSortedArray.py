# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def removeDuplicates_1(self, nums: list[int]) -> int:
        seen = set()
        lastUnique = 0
        for num in nums:
            if num not in seen:
                seen.add(num)
                nums[lastUnique] = num
                lastUnique += 1

        return len(seen)

    # Time complexity: O(n)
    # Space complexity: O(1)
    def removeDuplicates_2(self, nums: list[int]) -> int:
        l = 1
        for r in range(1, len(nums)):
            if nums[r] != nums[r - 1]:
                nums[l] = nums[r]
                l += 1

        return l
