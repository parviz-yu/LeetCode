# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def findDuplicate_extraSpace(self, nums: list[int]) -> int:
        seen = [False] * len(nums)
        for num in nums:
            if seen[num]:
                return num

            seen[num] = True

        return -1

    # Time complexity: O(n)
    # Space complexity: O(1)
    def findDuplicate_floydAlgo(self, nums: list[int]) -> int:
        slow = fast = nums[0]  # nums[0] is the head of linked list
        isEqual = False
        while not isEqual:
            slow = nums[slow]
            fast = nums[nums[fast]]

            if slow == fast:
                isEqual = True

        slow = nums[0]
        while slow != fast:
            slow = nums[slow]
            fast = nums[fast]

        return slow
