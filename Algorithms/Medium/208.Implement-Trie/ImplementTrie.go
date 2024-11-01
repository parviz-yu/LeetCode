package medium

type Node struct {
	Children   [26]*Node
	IsTerminal bool
}

type Trie struct {
	Root *Node
}

func Constructor() Trie {
	return Trie{Root: &Node{}}
}

// Time: O(n)
func (this *Trie) Insert(word string) {
	node := this.Root
	for _, c := range word {
		idx := c - 'a'
		if node.Children[idx] == nil {
			node.Children[idx] = &Node{}
		}
		node = node.Children[idx]
	}
	node.IsTerminal = true
}

// Time: O(n)
func (this *Trie) Search(word string) bool {
	node := this.Root
	for _, c := range word {
		idx := c - 'a'
		if node.Children[idx] == nil {
			return false
		}
		node = node.Children[idx]
	}

	return node.IsTerminal
}

// Time: O(n)
func (this *Trie) StartsWith(prefix string) bool {
	node := this.Root
	for _, c := range prefix {
		idx := c - 'a'
		if node.Children[idx] == nil {
			return false
		}

		node = node.Children[idx]
	}

	return true
}
