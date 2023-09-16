# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def longestConsecutive(self, nums: list[int]) -> int:
        hash_set = set(nums)
        max_length = 0

        for num in hash_set:
            if (num - 1) not in hash_set:
                next_num = num + 1
                while next_num in hash_set:
                    next_num += 1

                max_length = max(max_length, next_num - num)

        return max_length