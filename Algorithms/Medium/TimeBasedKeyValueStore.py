class TimeMap:
    def __init__(self):
        self.hash_table = {}

    def set(self, key: str, value: str, timestamp: int) -> None:
        self.hash_table.setdefault(key, []).append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        val = self.hash_table.get(key, None)
        if val is None:
            return ""

        l, r = 0, len(val) - 1
        while l < r:
            m = (l + r + 1) // 2
            if timestamp >= val[m][0]:
                l = m
            else:
                r = m - 1

        if val[l][0] <= timestamp:
            return val[l][1]

        return ""
