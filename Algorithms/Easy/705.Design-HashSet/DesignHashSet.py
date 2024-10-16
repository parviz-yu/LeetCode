class Node:
    def __init__(self, val, next=None):
        self.val = val
        self.next = next


class MyHashSet:
    def __init__(self):
        self.array = [None] * 10000
        self.length = len(self.array)

    def calc_hash_value(self, key: int) -> int:
        return key % self.length

    def add(self, key: int) -> None:
        if self.contains(key):
            return None

        hash_val = self.calc_hash_value(key)
        if self.array[hash_val] is None:
            self.array[hash_val] = Node(key)
            return

        curr = self.array[hash_val]
        while curr.next:
            curr = curr.next

        curr.next = Node(key)

    def remove(self, key: int) -> None:
        if not self.contains(key):
            return None

        hash_val = self.calc_hash_value(key)
        if self.array[hash_val].val == key:
            self.array[hash_val] = self.array[hash_val].next
            return

        prev, curr = None, self.array[hash_val]
        while curr.val != key:
            prev = curr
            curr = curr.next

        prev.next = curr.next

    def contains(self, key: int) -> bool:
        curr = self.array[self.calc_hash_value(key)]
        while curr:
            if curr.val == key:
                return True

            curr = curr.next

        return False
