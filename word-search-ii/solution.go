package wordsearchii

func findWords(board [][]byte, words []string) []string {
	wordsFound := []string{}
	var dfs func(position, *trieNode, []byte)
	dfs = func(pos position, tn *trieNode, w []byte) {
		if pos.row < 0 || pos.row >= len(board) || pos.col < 0 || pos.col >= len(board[0]) {
			return
		}

		if board[pos.row][pos.col] == '$' {
			return
		}

		char := board[pos.row][pos.col]
		board[pos.row][pos.col] = '$'
		defer func() {
			board[pos.row][pos.col] = char
		}()

		charNode, found := tn.children[char]
		if !found {
			return
		}

		w = append(w, char)

		if charNode.end > 0 {
			charNode.end -= 1
			wordsFound = append(wordsFound, string(w))
		}

		for _, nextPosition := range nextPositions(pos) {
			dfs(nextPosition, charNode, w)
		}
	}

	trie := newTrie()
	for _, word := range words {
		insert(trie, word)
	}
	for r, row := range board {
		for c := range row {
			dfs(position{r, c}, trie, nil)
		}
	}
	return wordsFound
}

type position struct {
	row, col int
}

func nextPositions(p position) []position {
	return []position{
		{p.row - 1, p.col},
		{p.row, p.col + 1},
		{p.row + 1, p.col},
		{p.row, p.col - 1},
	}
}

type trieNode struct {
	char     byte
	children map[byte]*trieNode
	end      int
}

func newTrieNode(char byte) *trieNode {
	return &trieNode{
		char:     char,
		children: make(map[byte]*trieNode),
		end:      0,
	}
}

type trie = *trieNode

func newTrie() trie { return newTrieNode('/') }

func insert(t trie, word string) {
	curr := t
	for i := 0; i < len(word); i++ {
		if _, found := curr.children[word[i]]; !found {
			curr.children[word[i]] = newTrieNode(word[i])
		}
		curr = curr.children[word[i]]
	}
	curr.end += 1
}
