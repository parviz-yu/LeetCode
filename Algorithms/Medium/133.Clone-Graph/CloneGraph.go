package medium

type Node struct {
	Val       int
	Neighbors []*Node
}

// Time: O(v + e)
// Space: O(v)
func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	visited := make(map[*Node]*Node)
	stack := make([]*Node, 0)
	visited[node] = &Node{Val: node.Val}
	stack = append(stack, node)
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, n := range v.Neighbors {
			if _, ok := visited[n]; !ok {
				visited[n] = &Node{Val: n.Val}
				stack = append(stack, n)
			}

			visited[v].Neighbors = append(visited[v].Neighbors, visited[n])
		}
	}

	return visited[node]
}
