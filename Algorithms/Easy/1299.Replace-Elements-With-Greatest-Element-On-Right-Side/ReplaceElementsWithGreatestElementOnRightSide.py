# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def replaceElements(self, arr: list[int]) -> list[int]:
        curr_max = -1

        for i in range(len(arr) - 1, -1, -1):
            new_max = max(curr_max, arr[i])
            arr[i] = curr_max
            curr_max = new_max

        return arr
