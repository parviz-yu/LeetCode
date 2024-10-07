package easy

import "github.com/parviz-yu/LeetCode/Algorithms/helpers"

type MyStack struct {
	queue *helpers.Queue
}

func NewMyStack() MyStack {
	return MyStack{
		queue: &helpers.Queue{},
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
