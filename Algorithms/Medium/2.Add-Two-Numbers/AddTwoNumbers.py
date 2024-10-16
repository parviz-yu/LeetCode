from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = curr = ListNode()

        add_next = 0
        while l1 or l2 or add_next:
            val_1 = 0
            if l1:
                val_1 = l1.val
                l1 = l1.next

            val_2 = 0
            if l2:
                val_2 = l2.val
                l2 = l2.next

            add_next, curr_sum = divmod(val_1 + val_2 + add_next, 10)
            curr.next = ListNode(curr_sum)
            curr = curr.next

        return dummy.next