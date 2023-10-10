from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


# Time complexity: O(m + n)
class Solution:
    # Space complexity: O(m)
    def getIntersectionNode_v1(
        self, headA: ListNode, headB: ListNode
    ) -> Optional[ListNode]:
        seen = set()
        while headA:
            if headA not in seen:
                seen.add(headA)

            headA = headA.next

        while headB:
            if headB in seen:
                return headB

            headB = headB.next

        return None

    # Space complexity: O(1)
    def getIntersectionNode_v2(
        self, headA: ListNode, headB: ListNode
    ) -> Optional[ListNode]:
        lenA = self.length(headA)
        lenB = self.length(headB)

        if lenA > lenB:
            headA = self.skipNodes(headA, lenA - lenB)
        elif lenB > lenA:
            headB = self.skipNodes(headB, lenB - lenA)

        while headA and headB:
            if headA == headB:
                return headA

            headA = headA.next
            headB = headB.next

        return None

    def length(self, head: ListNode) -> int:
        curr = head
        count = 0
        while curr:
            count += 1
            curr = curr.next

        return count

    def skipNodes(self, head: ListNode, n: int) -> None:
        for _ in range(n):
            head = head.next

        return head
