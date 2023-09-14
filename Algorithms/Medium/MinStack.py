# Time complexity: O(1)
class MinStack:
    def __init__(self):
        self.stack = []
        self.min_stack = []

    def push(self, val: int) -> None:
        if len(self.min_stack) == 0:
            self.min_stack.append(val)
        else:
            self.min_stack.append(min(self.min_stack[-1], val))

        self.stack.append(val)
        return None

    def pop(self) -> None:
        self.stack.pop()
        self.min_stack.pop()
        return None

    def top(self) -> int:
        return self.stack[-1]

    def getMin(self) -> int:
        return self.min_stack[-1] 
