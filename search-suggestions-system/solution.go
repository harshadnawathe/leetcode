package searchsuggestionssystem

type trieNode struct {
	children [26]*trieNode
	indexes  []int
	b        byte
}

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := &trieNode{b: '$'}
	for index, product := range products {
		node := trie
		for i := 0; i < len(product); i++ {
			c := product[i]
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &trieNode{b: c}
			}
			node = node.children[c-'a']
			node.indexes = append(node.indexes, index)
			for j := len(node.indexes) - 2; j >= 0 && products[node.indexes[j]] > products[node.indexes[j+1]]; j-- {
				node.indexes[j], node.indexes[j+1] = node.indexes[j+1], node.indexes[j]
			}
			if len(node.indexes) > 3 {
				node.indexes = node.indexes[:3]
			}
		}
	}

	result := [][]string{}
	for i := 0; i < len(searchWord); i++ {
		c := searchWord[i]
		trie = trie.children[c-'a']
		if trie == nil {
			break
		}
		suggestions := []string{}
		for _, i := range trie.indexes {
			suggestions = append(suggestions, products[i])
		}
		result = append(result, suggestions)
	}

	if len(result) < len(searchWord) {
		for i := 0; len(searchWord) != len(result); i++ {
			result = append(result, []string{})
		}
	}

	return result
}
