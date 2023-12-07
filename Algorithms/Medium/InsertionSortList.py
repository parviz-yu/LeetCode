from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# Time complexity: O(n^2)
# Space complexity: O(1)
class Solution:
    def insertionSortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode()
        curr = head

        while curr:
            prev = dummy
            while prev.next and prev.next.val <= curr.val:
                prev = prev.next

            nxt = curr.next
            curr.next = prev.next
            prev.next = curr
            curr = nxt

        return dummy.next
