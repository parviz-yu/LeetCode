package helpers

type ListNode struct {
	Val  int
	Next *ListNode
}

type Queue struct {
	Head *ListNode
	Tail *ListNode
	Size int
}

func (q *Queue) Enqueue(x int) {
	node := &ListNode{Val: x}
	q.Size++
	if q.Tail == nil {
		q.Tail = node
		q.Head = node
		return
	}

	q.Tail.Next = node
	q.Tail = node
}

func (q *Queue) Dequeue() int {
	node := q.Head
	q.Head = node.Next
	q.Size--
	if q.Size == 0 {
		q.Tail = nil
	}

	return node.Val
}
