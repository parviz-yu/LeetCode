package easy

type MyQueue struct {
	enqueue []int
	dequeue []int
}

func NewMyQueue() MyQueue {
	return MyQueue{
		enqueue: make([]int, 0),
		dequeue: make([]int, 0),
	}
}

// Time: O(1)
func (q *MyQueue) Push(x int) {
	q.enqueue = append(q.enqueue, x)
}

// Time: O(1)
func (q *MyQueue) Pop() int {
	if len(q.dequeue) == 0 {
		q.move()
	}

	value := q.dequeue[len(q.dequeue)-1]
	q.dequeue = q.dequeue[:len(q.dequeue)-1]
	return value
}

// Time: O(1)
func (q *MyQueue) Peek() int {
	if len(q.dequeue) == 0 {
		q.move()
	}

	return q.dequeue[len(q.dequeue)-1]
}

// Time: O(1)
func (q *MyQueue) Empty() bool {
	return len(q.enqueue) == 0 && len(q.dequeue) == 0
}

// Time: O(n)
func (q *MyQueue) move() {
	for i := len(q.enqueue) - 1; i >= 0; i-- {
		q.dequeue = append(q.dequeue, q.enqueue[i])
	}
	q.enqueue = q.enqueue[:0]
}
