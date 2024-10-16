class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        letters_cnt = dict()
        for char in s:
            letters_cnt[char] = letters_cnt.get(char, 0) + 1

        for char in t:
            if letters_cnt.get(char) == None:
                return False

            if letters_cnt.get(char) == 0:
                return False

            letters_cnt[char] -= 1

        for key in letters_cnt.keys():
            if letters_cnt[key] != 0:
                return False

        return True