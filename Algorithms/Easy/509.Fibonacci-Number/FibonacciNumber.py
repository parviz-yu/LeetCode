class Solution:
    # Time complexity: O(2^n)
    # Space complexity: O(n)
    def fib_1(self, n: int) -> int:
        if n < 2:
            return n

        return self.fib_1(n - 1) + self.fib_1(n - 2)

    # Time complexity: O(n)
    # Space complexity: O(n)
    def fib_2(self, n: int) -> int:
        arr = [0, 1]

        for i in range(2, n + 1):
            arr.append(arr[i - 1] + arr[i - 2])

        print(arr)

        return arr[n]

    # Time complexity: O(n)
    # Space complexity: O(1)
    def fib_3(self, n: int) -> int:
        if n <= 1:
            return n

        a, b = 0, 1
        for i in range(2, n + 1):
            a, b = b, a + b

        return b
