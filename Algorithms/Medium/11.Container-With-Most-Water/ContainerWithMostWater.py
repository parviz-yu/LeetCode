# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def maxArea(self, height: list[int]) -> int:
        max_area = 0
        l, r = 0, len(height) - 1
        while l < r:
            area = min(height[l], height[r]) * (r - l)
            if area > max_area:
                max_area = area
            
            if height[l] < height[r]:
                l += 1
            else:
                r -= 1
        
        return max_area