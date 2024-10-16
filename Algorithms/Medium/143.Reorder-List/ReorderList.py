from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        stack = []
        slow, fast = head, head

        # getting the middle of the linked list
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next

        end = slow
        while slow:
            stack.append(slow)
            slow = slow.next

        while head != end:
            nxt = head.next
            node = stack.pop()
            head.next = node
            node.next = nxt
            head = head.next.next

        end.next = None