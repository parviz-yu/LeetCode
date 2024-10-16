package easy

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

type MyStack struct {
	queue *Queue
}

func NewMyStack() MyStack {
	return MyStack{
		queue: &Queue{},
	}
}

// Time: O(1)
func (this *MyStack) Push(x int) {
	this.queue.Enqueue(x)
}

// Time: O(n)
func (this *MyStack) Pop() int {
	for _ = range this.queue.Size - 1 {
		x := this.queue.Dequeue()
		this.queue.Enqueue(x)
	}

	return this.queue.Dequeue()
}

// Time: O(1)
func (this *MyStack) Top() int {
	return this.queue.Tail.Val
}

// Time: O(1)
func (this *MyStack) Empty() bool {
	return this.queue.Size == 0
}
