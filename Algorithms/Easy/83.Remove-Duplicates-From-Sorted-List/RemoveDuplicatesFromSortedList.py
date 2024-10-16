from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(-101, head)
        curr, nxt = dummy, head

        while nxt:
            if curr.val == nxt.val:
                curr.next = nxt.next
            else:
                curr = curr.next
            nxt = nxt.next

        return dummy.next
