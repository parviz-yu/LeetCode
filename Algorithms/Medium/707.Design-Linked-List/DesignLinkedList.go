package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(val int, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: next}
}

type MyLinkedList struct {
	Head *ListNode
	Tail *ListNode
	Len  int
}

func NewLinkedList() MyLinkedList {
	return MyLinkedList{}
}

// Time complexity: O(n)
func (this *MyLinkedList) Get(index int) int {
	if index > this.Len-1 {
		return -1
	}

	curr := this.Head
	for index > 0 {
		curr = curr.Next
		index -= 1
	}

	return curr.Val
}

// Time complexity: O(1)
func (this *MyLinkedList) AddAtHead(val int) {
	this.Len += 1
	this.Head = NewNode(val, this.Head)
	if this.Tail == nil {
		this.Tail = this.Head
	}
}

// Time complexity: O(1)
func (this *MyLinkedList) AddAtTail(val int) {
	if this.Tail == nil && this.Head == nil {
		this.AddAtHead(val)
		return
	}

	this.Len += 1
	this.Tail.Next = NewNode(val, nil)
	this.Tail = this.Tail.Next
}

// Time complexity: O(n)
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Len {
		return
	}

	if index == this.Len {
		this.AddAtTail(val)
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	this.Len += 1
	curr := NewNode(0, this.Head)
	for index > 0 {
		curr = curr.Next
		index -= 1
	}

	nxt := curr.Next
	curr.Next = NewNode(val, nxt)
}

// Time complexity: O(n)
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.Len-1 {
		return
	}

	this.Len -= 1
	if index == 0 {
		this.Head = this.Head.Next
		return
	}

	prev := NewNode(0, this.Head)
	for index > 0 {
		prev = prev.Next
		index -= 1
	}

	if prev.Next == this.Tail {
		this.Tail = prev
	}

	prev.Next = prev.Next.Next
}
