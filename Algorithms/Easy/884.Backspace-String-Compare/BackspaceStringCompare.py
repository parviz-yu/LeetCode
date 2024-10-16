class Solution:
    # Time: O(m + n)
    # Space: O(m + n)
    def backspaceCompare_1(self, s: str, t: str) -> bool:
        s_stack = self.removeBackspace(s)
        t_stack = self.removeBackspace(t)

        if len(s_stack) != len(t_stack):
            return False

        for i in range(len(s_stack)):
            if s_stack[i] != t_stack[i]:
                return False

        return True

    def removeBackspace(self, s: str) -> list[str]:
        stack = []
        for i in range(len(s)):
            if s[i] == "#":
                if len(stack) > 0:
                    stack.pop()
                continue

            stack.append(s[i])

        return stack

    # Time: O(m + n)
    # Space: O(1)
    def backspaceCompare_2(self, s: str, t: str) -> bool:
        i, j = len(s) - 1, len(t) - 1

        while True:
            skip = 0
            while i >= 0 and (s[i] == "#" or skip > 0):
                skip += 1 if s[i] == "#" else -1
                i -= 1

            skip = 0
            while j >= 0 and (t[j] == "#" or skip > 0):
                skip += 1 if t[j] == "#" else -1
                j -= 1

            if i >= 0 and j >=0 and s[i] == t[j]:
                i -= 1
                j -= 1
            else:
                return i == -1 and j == -1