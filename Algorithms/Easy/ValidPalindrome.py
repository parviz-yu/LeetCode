# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def isPalindrome(self, s: str) -> bool:
        buf = []
        for c in s.lower():
            if c.isdigit() or c.isalpha():
                buf.append(c)

        s = "".join(buf)
        l, r = 0, len(s) - 1
        while l <= r:
            if s[l] != s[r]:
                return False
            
            l += 1
            r -= 1
        
        return True