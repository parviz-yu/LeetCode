from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        less, more = ListNode(), ListNode()
        l, m = less, more
        while head:
            if head.val < x:
                l.next = head
                l = l.next
            else:
                m.next = head
                m = m.next

            head = head.next

        l.next = m.next = None
        l.next = more.next
        return less.next
