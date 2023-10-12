class Node:
    def __init__(self, val, next=None):
        self.val = val
        self.next = next


class MyLinkedList:
    def __init__(self):
        self.head = None
        self.tail = None
        self.len = 0

    # Time complexity: O(n)
    def get(self, index: int) -> int:
        if index > self.len - 1:
            return -1

        curr = self.head
        while index > 0:
            curr = curr.next
            index -= 1

        return curr.val

    # Time complexity: O(1)
    def addAtHead(self, val: int) -> None:
        self.len += 1
        self.head = Node(val, self.head)
        if self.tail is None:
            self.tail = self.head

    # Time complexity: O(1)
    def addAtTail(self, val: int) -> None:
        if self.head is None and self.tail is None:
            return self.addAtHead(val)

        self.len += 1
        new_node = Node(val)
        self.tail.next = new_node
        self.tail = new_node

    # Time complexity: O(n)
    def addAtIndex(self, index: int, val: int) -> None:
        if index > self.len:
            return

        if index == self.len:
            return self.addAtTail(val)

        if index == 0:
            return self.addAtHead(val)

        self.len += 1
        prev = Node(0, self.head)
        while index > 0:
            prev = prev.next
            index -= 1

        nxt = prev.next
        prev.next = Node(val, nxt)

    # Time complexity: O(n)
    def deleteAtIndex(self, index: int) -> None:
        if index > self.len - 1:
            return

        self.len -= 1
        if index == 0:
            self.head = self.head.next
            return

        prev = Node(0, self.head)
        while index > 0:
            prev = prev.next
            index -= 1

        if prev.next == self.tail:
            self.tail = prev

        prev.next = prev.next.next
