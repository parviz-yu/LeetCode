from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# Time complexity: O(n)
# Space complexity: O(1)
class Solution:
    def removeNthFromEnd_withCouter(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        count = 0
        curr = head
        while curr:
            count += 1
            curr = curr.next

        if count == n:
            head = head.next
        else:    
            curr = head
            prev = None
            while curr:
                if count == n:
                    prev.next = curr.next

                count -= 1
                prev = curr
                curr = curr.next

        return head
    

    def removeNthFromEnd_onePass(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        fast, slow = head, head
        for _ in range(n):
            fast = fast.next

        if fast is None:
            return head.next

        while fast.next:
            fast, slow = fast.next, slow.next

        slow.next = slow.next.next

        return head