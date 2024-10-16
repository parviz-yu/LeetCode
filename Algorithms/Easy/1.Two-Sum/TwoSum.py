# Time complexity O(n)
# Space complexity O(n)
class Solution(object):
    def twoSum(self, nums, target):
        subtr = {}

        for i in range(len(nums)):
            diff = target - nums[i]
            if diff in subtr:
                return [i, subtr[diff]]
            
            subtr[nums[i]] = i