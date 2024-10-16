package medium

type NodeRandom struct {
	Val    int
	Next   *NodeRandom
	Random *NodeRandom
}

// Time complexity: O(n)
// Space complexity: O(n)
func copyRandomList(head *NodeRandom) *NodeRandom {
	hashMap := make(map[*NodeRandom]*NodeRandom, 0)

	curr := head
	for curr != nil {
		hashMap[curr] = &NodeRandom{curr.Val, nil, nil}
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
