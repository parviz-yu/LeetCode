from typing import Optional


class Node:
    def __init__(self, x: int, next: "Node" = None, random: "Node" = None):
        self.val = int(x)
        self.next = next
        self.random = random


# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def copyRandomList(self, head: Optional[Node]) -> Optional[Node]:
        hash_table = {}
        curr = head

        while curr:
            hash_table[curr] = Node(curr.val)
            curr = curr.next

        copy = hash_table.get(head, None)
        while head:
            nxt = hash_table.get(head.next, None)
            rnd = hash_table.get(head.random, None)
            hash_table[head].next = nxt
            hash_table[head].random = rnd
            head = head.next

        return copy
