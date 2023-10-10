package medium

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// Time complexity: O(n)
// Space complexity: O(n)
func copyRandomList(head *Node) *Node {
	hashMap := make(map[*Node]*Node, 0)

	curr := head
	for curr != nil {
		hashMap[curr] = &Node{curr.Val, nil, nil}
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		hashMap[curr].Next = hashMap[curr.Next]
		hashMap[curr].Random = hashMap[curr.Random]
		curr = curr.Next
	}

	return hashMap[head]
}
