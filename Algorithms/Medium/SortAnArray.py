# Time complexity: O(nlogn)
# Space complexity: O(n)
class Solution:
    def sortArray_1(self, nums: list[int]) -> list[int]:
        def merge_sort(nums: list[int], l: int, r: int, temp: list[int]) -> None:
            if l < r:
                m = (l + r) // 2
                merge_sort(nums, l, m, temp)
                merge_sort(nums, m + 1, r, temp)
                merge(nums, l, m + 1, r, temp)

        def merge(
            nums: list[int],
            a_first_idx: int,
            b_first_idx: int,
            b_last_idx: int,
            temp_list: list[int],
        ):
            i, j = a_first_idx, b_first_idx
            k = i

            while i < b_first_idx and j <= b_last_idx:
                if nums[i] <= nums[j]:
                    temp_list[k] = nums[i]
                    i += 1
                else:
                    temp_list[k] = nums[j]
                    j += 1
                k += 1

            while i < b_first_idx:
                temp_list[k] = nums[i]
                i += 1
                k += 1

            while j <= b_last_idx:
                temp_list[k] = nums[j]
                j += 1
                k += 1

            for i in range(a_first_idx, b_last_idx + 1):
                nums[i] = temp_list[i]

        temp_list = [0] * len(nums)
        merge_sort(nums, 0, len(nums) - 1, temp_list)
        return nums

    def sortArray_2(self, nums: list[int]) -> list[int]:

        def quick_sort(nums, low, high):
            if low < high:
                d = partitiion(nums, low, high)
                quick_sort(nums, low, d - 1)
                quick_sort(nums, d + 1, high)

        def partitiion(nums, low, high):
            pivot = nums[high]
            firstLarger = low
            for i in range(low, high):
                if nums[i] < pivot:
                    nums[i], nums[firstLarger] = nums[firstLarger], nums[i]
                    firstLarger += 1

            nums[firstLarger], nums[high] = nums[high], nums[firstLarger]
            return firstLarger

        quick_sort(nums, 0, len(nums) - 1)
        return nums
