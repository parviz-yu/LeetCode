# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def findDisappearedNumbers_1(self, nums: list[int]) -> list[int]:
        counter = [0] * (len(nums) + 1)
        for num in nums:
            counter[num] += 1

        ans = []
        for i in range(1, len(counter)):
            if counter[i] == 0:
                ans.append(i)

        return ans

    # Time complexity: O(n)
    # Space complexity: O(1)
    def findDisappearedNumbers_2(self, nums: list[int]) -> list[int]:
        ans = []
        n = len(nums)

        for i in range(n):
            idx = abs(nums[i]) - 1
            if nums[idx] > 0:
                nums[idx] *= -1

        for i in range(n):
            if nums[i] > 0:
                ans.append(i + 1)

        return ans
