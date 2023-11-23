# Time complexity: O(m + n)
# Space complexity: O(m)
class Solution:
    def merge(self, nums1: list[int], m: int, nums2: list[int], n: int) -> None:
        temp = nums1[:]

        i, j = 0, 0
        k = 0
        while i < m and j < n:
            if temp[i] < nums2[j]:
                nums1[k] = temp[i]
                i += 1
            else:
                nums1[k] = nums2[j]
                j += 1
            k += 1

        while i < m:
            nums1[k] = temp[i]
            i += 1
            k += 1

        while j < n:
            nums1[k] = nums2[j]
            j += 1
            k += 1
