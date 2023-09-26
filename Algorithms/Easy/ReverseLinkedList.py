from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        reversed_head = None       

        while head:
            curr = ListNode(head.val)
            curr.next = reversed_head
            reversed_head = curr
            head = head.next

        return reversed_head