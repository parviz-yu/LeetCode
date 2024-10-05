package medium

type LRUNode struct {
	Key   int
	Value int
	Next  *LRUNode
	Prev  *LRUNode
}

type LRUCache struct {
	Capacity int
	Data     map[int]*LRUNode
	Head     *LRUNode
	Tail     *LRUNode
}

func Constructor(capacity int) LRUCache {
	data := make(map[int]*LRUNode, capacity)
	head := &LRUNode{}
	tail := &LRUNode{}

	head.Next = tail
	tail.Prev = head

	return LRUCache{
		Capacity: capacity,
		Data:     data,
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) remove(node *LRUNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) add2Head(node *LRUNode) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Data[key]
	if !ok {
		return -1
	}

	this.remove(node)
	this.add2Head(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Data[key]
	if ok {
		this.remove(node)
		this.add2Head(node)
		node.Value = value
		return
	}

	if len(this.Data) == this.Capacity {
		node = this.Tail.Prev
		this.remove(node)
		delete(this.Data, node.Key)
	}

	node = &LRUNode{
		Key:   key,
		Value: value,
	}

	this.add2Head(node)
	this.Data[key] = node
	return
}
