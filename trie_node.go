package main

type TrieNode struct {
	Word     string
	Children map[byte]*TrieNode
}

func NewTrieNode() *TrieNode {
	n := &TrieNode{
		Children: map[byte]*TrieNode{},
	}
	return n
}

func (n *TrieNode) Insert(word string) {
	current := n
	bytes := []byte(word)
	for _, b := range bytes {
		if _, ok := current.Children[b]; !ok {
			current.Children[b] = NewTrieNode()
		}
		current = current.Children[b]
	}
	current.Word = word
}
