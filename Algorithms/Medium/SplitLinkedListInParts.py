from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# Time complexity: O(k + n)
# Space complexity: O(1)
class Solution:
    def splitListToParts(
        self, head: Optional[ListNode], k: int
    ) -> list[Optional[ListNode]]:
        result = [None] * k

        length = self.listLen(head)
        one_more = length % k

        curr = head
        for i in range(k):
            skip = length // k - 1
            if one_more > 0:
                skip += 1
                one_more -= 1

            for _ in range(skip):
                curr = curr.next

            if head:
                result[i] = head
                head = curr.next
                curr.next = None
                curr = head

        return result

    def listLen(self, head: Optional[ListNode]) -> int:
        n, curr = 0, head
        while curr:
            n += 1
            curr = curr.next

        return n
