# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def canPlaceFlowers(self, flowerbed: list[int], n: int) -> bool:
        if n == 0:
            return True

        flowerbed = [0] + flowerbed + [0]

        for i in range(1, len(flowerbed) - 1):
            if flowerbed[i] + flowerbed[i - 1] + flowerbed[i + 1] == 0:
                flowerbed[i] = 1
                n -= 1

            if n == 0:
                return True

        return False


# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def canPlaceFlowers(self, flowerbed: list[int], n: int) -> bool:
        max_num = 0
        for i in range(len(flowerbed)):
            if flowerbed[i] == 0:
                is_left_empty = (i == 0) or (flowerbed[i - 1] == 0)
                is_right_empty = (i == len(flowerbed) - 1) or (flowerbed[i + 1] == 0)
                if is_left_empty and is_right_empty:
                    flowerbed[i] = 1
                    max_num += 1

        return max_num >= n
