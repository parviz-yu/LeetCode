# Time complexity - O(n)
# Space complexity - O(n)
class Solution:
    def containsDuplicate(self, nums: list[int]) -> bool:
        exists = set()
        for num in nums:
            if num not in exists:
                exists.add(num)
            else:
                return True

        return False