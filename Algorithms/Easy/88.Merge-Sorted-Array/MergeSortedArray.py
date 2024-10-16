# Time complexity: O(m + n)
class Solution:
    # Space complexity: O(m)
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

    # Space complexity: O(1)
    def merge_1(self, nums1: list[int], m: int, nums2: list[int], n: int) -> None:
        last = n + m - 1
        while m > 0 and n > 0:
            if nums1[m - 1] > nums2[n - 1]:
                nums1[last] = nums1[m - 1]
                m -= 1
            else:
                nums1[last] = nums2[n - 1]
                n -= 1
            last -= 1

        while n > 0:
            nums1[last] = nums2[n - 1]
            n -= 1
            last -= 1
