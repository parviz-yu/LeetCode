package easy

type Node struct {
	Key  int
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Get(key int) int {
	curr := ll.Head
	for curr != nil {
		if curr.Key == key {
			return curr.Val
		}

		curr = curr.Next
	}

	return -1
}

func (ll *LinkedList) Put(key int, val int) {
	curr := ll.Head
	for curr != nil {
		if curr.Key == key {
			curr.Val = val
			return
		}

		curr = curr.Next
	}

	ll.Head = &Node{
		Key:  key,
		Val:  val,
		Next: ll.Head,
	}
}

func (ll *LinkedList) Remove(key int) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Key == key {
		ll.Head = ll.Head.Next
		return
	}

	curr := ll.Head
	for curr.Next != nil {
		if curr.Next.Key == key {
			curr.Next = curr.Next.Next
			return
		}

		curr = curr.Next
	}
}

type MyHashMap struct {
	n       int
	buckets []*LinkedList
}

func NewMyHashMap() MyHashMap {
	n := 1021
	buckets := make([]*LinkedList, n)

	for i := range buckets {
		buckets[i] = &LinkedList{}
	}

	return MyHashMap{
		n:       n,
		buckets: buckets,
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % this.n
}

func (this *MyHashMap) Put(key int, value int) {
	this.buckets[this.hash(key)].Put(key, value)
}

func (this *MyHashMap) Get(key int) int {
	return this.buckets[this.hash(key)].Get(key)
}

func (this *MyHashMap) Remove(key int) {
	this.buckets[this.hash(key)].Remove(key)
}
