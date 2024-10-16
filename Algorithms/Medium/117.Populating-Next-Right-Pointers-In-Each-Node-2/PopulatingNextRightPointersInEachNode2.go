package medium

import "container/list"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// Time: O(n)
// Space: O(n)
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {

		var prev *Node
		for range queue.Len() {
			node := queue.Remove(queue.Front()).(*Node)

			if prev != nil {
				prev.Next = node
			}
			prev = node

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return root
}
