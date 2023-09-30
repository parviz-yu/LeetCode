from typing import Optional

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        dummy = ListNode()
        dummy.next = head
        slow, fast = dummy, head

        while slow and fast and fast.next:
            if slow == fast:
                return True
            slow = slow.next
            fast = fast.next.next

        return False