class Solution:
    # Time complexity: O(n)
    # Space complexity: O(n)
    def majorityElement_1(self, nums: list[int]) -> int:
        counter = {}
        for num in nums:
            counter[num] = counter.get(num, 0) + 1

        maxEl, maxCount = 0, 0
        for elem, count in counter.items():
            if count > maxCount:
                maxEl = elem
                maxCount = count

        return maxEl

    # Time complexity: O(n)
    # Space complexity: O(1)
    # Boyer–Moore majority vote algorithm
    def majorityElement_2(self, nums: list[int]) -> int:
        counter, res = 0, 0
        for num in nums:
            if counter == 0:
                res = num
                counter = 1
            elif num == res:
                counter += 1
            else:
                counter -= 1

        return res
