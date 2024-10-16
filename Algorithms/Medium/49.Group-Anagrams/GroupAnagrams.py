# Time complexity: O(mnlogn)
#   n - length of string
#   m - length of array
# Space complexity: O(n)
class Solution:
    def groupAnagrams(self, strs: list[str]) -> list[list[str]]:
        hash_map = {}

        for s in strs:
            s_sorted = "".join(sorted(s))
            if s_sorted not in hash_map:
                hash_map[s_sorted] = []

            hash_map[s_sorted].append(s)

        return list(hash_map.values())