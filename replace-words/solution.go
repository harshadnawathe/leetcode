package replacewords

import (
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	var b strings.Builder

	trie := newTrie(dictionary)

	node := trie
	for i := 0; i < len(sentence); i++ {
		char := sentence[i]

		if char == ' ' || node == nil || !node.IsEnd {
			b.WriteByte(char)
		}

		if char == ' ' {
			node = trie
		} else if node != nil && !node.IsEnd {
			node = node.CharTrieNode[char]
		}
	}

	return b.String()
}

type trieNode struct {
	CharTrieNode map[byte]*trieNode
	Char         byte
	IsEnd        bool
}

func newTrieNode(char byte) *trieNode {
	return &trieNode{
		make(map[byte]*trieNode),
		char,
		false,
	}
}

func newTrie(words []string) *trieNode {
	root := newTrieNode('$')

	for _, word := range words {
		node := root

		for i := range len(word) {
			var next *trieNode
			if charNode, ok := node.CharTrieNode[word[i]]; ok {
				next = charNode
			} else {
				next = newTrieNode(word[i])
				node.CharTrieNode[word[i]] = next
			}
			node = next
		}
		node.IsEnd = true
	}

	return root
}
