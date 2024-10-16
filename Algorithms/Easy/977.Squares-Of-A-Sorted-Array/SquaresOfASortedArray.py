# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def sortedSquares_1(self, nums: list[int]) -> list[int]:
        res = []

        # get the last elem that <= 0 with the binary search
        n = len(nums)
        l, r = 0, n - 1
        while l < r:
            m = (l + r + 1) // 2
            if nums[m] <= 0:
                l = m
            else:
                r = m - 1

        r = l + 1

        for i in range(len(nums)):
            nums[i] = nums[i] ** 2

        # place squares in the right order
        while l > -1 and r < n:
            if nums[l] < nums[r]:
                res.append(nums[l])
                l -= 1
            else:
                res.append(nums[r])
                r += 1

        while l > -1:
            res.append(nums[l])
            l -= 1

        while r < n:
            res.append(nums[r])
            r += 1

        return res

    def sortedSquares_2(self, nums: list[int]) -> list[int]:
        n = len(nums)
        res = [0] * n
        l, r = 0, n - 1
        for i in range(n - 1, -1, -1):
            if abs(nums[l]) > abs(nums[r]):
                res[i] = nums[l] ** 2
                l += 1
            else:
                res[i] = nums[r] ** 2
                r -= 1

        return res
