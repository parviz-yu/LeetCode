class Solution:
    # Time complexity: O(n)
    # Space complexity: O(n)
    def rotate_1(self, nums: list[int], k: int) -> None:
        n = len(nums)
        rotated = nums[:]

        for i in range(n):
            j = (i + k) % n
            nums[j] = rotated[i]

    # Time complexity: O(n)
    # Space complexity: O(1)
    def rotate_2(self, nums: list[int], k: int) -> None:
        n = len(nums)
        k %= n

        self.rotate_array(nums, 0, n - k - 1)
        self.rotate_array(nums, n - k, n - 1)
        self.rotate_array(nums, 0, n - 1)

    def rotate_array(self, arr: list[int], l: int, r: int):
        while l < r:
            arr[l], arr[r] = arr[r], arr[l]
            l += 1
            r -= 1
