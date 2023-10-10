from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def pairSum(self, head: Optional[ListNode]) -> int:
        slow, fast = head, head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next

        prev = None
        while slow:
            nxt = slow.next
            slow.next = prev
            prev = slow
            slow = nxt

        max_sum = 0
        while head and prev:
            max_sum = max(max_sum, head.val + prev.val)
            head = head.next
            prev = prev.next

        return max_sum
