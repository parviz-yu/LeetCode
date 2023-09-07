# Time complexity: O(n^2)
# Space complexity: O(1)
class Solution:
    def threeSum(self, nums: list[int]) -> list[list[int]]:
        result = []
        nums.sort()

        n = len(nums)
        for i in range(n):
            if i > 0 and nums[i] == nums[i-1]:
                continue

            l, r = i + 1, n - 1
            while l < r:
                curr_sum = nums[i] + nums[l] + nums[r]
                if curr_sum < 0:
                    l += 1
                elif curr_sum > 0:
                    r -= 1
                else:
                    result.append([nums[i], nums[l], nums[r]])
                    while l < r and nums[l] == nums[l+1]:
                        l += 1
                    while l < r and nums[r] == nums[r-1]:
                        r -= 1
                    l += 1
                    r -= 1
        
        return result