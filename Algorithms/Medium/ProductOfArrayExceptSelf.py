class Solution:
    # Time complexity: O(n)
    # Space complexity: O(n)
    def productExceptSelf_v1(self, nums: list[int]) -> list[int]:
        n = len(nums)
        result = [1] * n 
        prefix, postfix = [1] * (n + 1), [1] * (n + 1)

        for i in range(n):
            prefix[i + 1] = prefix[i] * nums[i]

        for i in range(n - 1, -1, -1):
            postfix[i] = postfix[i + 1] * nums[i]

        for i in range(n):
            result[i] = prefix[i] * postfix[i + 1]

        return result
    

    # Time complexity: O(n)
    # Space complexity: O(1)
    def productExceptSelf_v2(self, nums: list[int]) -> list[int]:
        n = len(nums)
        result = [1] * n 
        
        product = 1
        for i in range(n - 1):
            product *= nums[i]
            result[i + 1] = product

        product = 1
        for i in range(n - 1, 0, -1):
            product *= nums[i]
            result[i - 1] *= product

        return result