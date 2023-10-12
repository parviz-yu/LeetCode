from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    # Time complexity: O(n)
    # Space complexity: O(n)
    def swapNodes_v1(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        arr = [0]
        while head:
            arr.append(head.val)
            head = head.next

        arr[k], arr[len(arr) - k] = arr[len(arr) - k], arr[k]
        head = dummy = ListNode()
        for i in range(1, len(arr)):
            dummy.next = ListNode(arr[i])
            dummy = dummy.next

        return head.next

    # Time complexity: O(n)
    # Space complexity: O(1)
    def swapNodes_v2(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        slow, fast = head, head
        for _ in range(k - 1):
            fast = fast.next

        curr = fast
        while fast.next:
            fast = fast.next
            slow = slow.next

        curr.val, slow.val = slow.val, curr.val

        return head
