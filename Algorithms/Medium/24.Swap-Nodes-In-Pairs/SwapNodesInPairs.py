from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(next=head)
        prev, curr = dummy, head

        while curr and curr.next:
            nxt = curr.next
            nxt_nxt = curr.next.next

            nxt.next = curr
            prev.next = nxt
            curr.next = nxt_nxt

            prev = curr
            curr = nxt_nxt

        return dummy.next
