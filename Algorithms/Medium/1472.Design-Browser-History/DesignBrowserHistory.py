class Node:
    def __init__(self, value: str, next=None, prev=None):
        self.value = value
        self.next = next
        self.prev = prev


class BrowserHistory:
    def __init__(self, homepage: str):
        self.curr_page = Node(homepage)

    # Time complexity: O(1)
    def visit(self, url: str) -> None:
        new_page = Node(url, prev=self.curr_page)
        self.curr_page.next = new_page
        self.curr_page = new_page

    # Time complexity: O(n)
    def back(self, steps: int) -> str:
        while self.curr_page.prev and steps > 0:
            self.curr_page = self.curr_page.prev
            steps -= 1

        return self.curr_page.value

    # Time complexity: O(n)
    def forward(self, steps: int) -> str:
        while self.curr_page.next and steps > 0:
            self.curr_page = self.curr_page.next
            steps -= 1

        return self.curr_page.value
