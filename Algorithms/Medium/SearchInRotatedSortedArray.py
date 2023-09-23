# Time comlexity: O(logn)
# Space complexity: O(1)
class Solution:
    def search_1(self, nums: list[int], target: int) -> int:
        l, r = 0, len(nums) - 1
        while l < r:
            mid = (l + r) // 2
            if nums[mid] < nums[0]:
                r = mid
            else:
                l = mid + 1

        r = len(nums) - 1
        if target >= nums[0]:
            l, r = 0, l
        while l <= r:
            mid = (l + r) // 2
            if target == nums[mid]:
                return mid
            if nums[mid] > target:
                r = mid - 1
            else:
                l = mid + 1

        return -1
    

    def search_2(self, nums: list[int], target: int) -> int:
        l, r = 0, len(nums) - 1
        
        while l <= r:
            mid = l + (r - l) // 2
            if nums[mid] == target:
                return mid

            if nums[mid] >= nums[l]:
                if target >= nums[l] and target < nums[mid]:
                    r = mid - 1
                else:
                    l = mid + 1
            else:  
                if target <= nums[r] and target > nums[mid]:
                    l = mid + 1
                else:
                    r = mid - 1
                    
        return -1