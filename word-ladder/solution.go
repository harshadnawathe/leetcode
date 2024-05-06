package wordladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	graph := makeAdjacencyList(wordList)

	if _, ok := graph[endWord]; !ok {
		return 0
	}

	type node struct {
		w string
		n int
	}

	q := []node{}

	for word := range graph {
		if diffBy1(beginWord, word) {
			q = append(q, node{word, 2})
		}
	}

	visited := make(map[string]struct{})

	for len(q) > 0 {
		currNode := q[0]
		q = q[1:]

		if currNode.w == endWord {
			return currNode.n
		}

		if _, ok := visited[currNode.w]; ok {
			continue
		}

		visited[currNode.w] = struct{}{}

		for _, nextWord := range graph[currNode.w] {
			q = append(q, node{nextWord, currNode.n + 1})
		}
	}

	return 0
}

func makeAdjacencyList(wordList []string) map[string][]string {
	g := make(map[string][]string)

	for _, word := range wordList {
		g[word] = []string{}
	}

	for i, word1 := range wordList {
		for _, word2 := range wordList[i+1:] {
			if diffBy1(word1, word2) {
				g[word1] = append(g[word1], word2)
				g[word2] = append(g[word2], word1)
			}
		}
	}

	return g
}

func diffBy1(w1, w2 string) bool {
	diffFound := false
	for i := range w1 {
		if w1[i] != w2[i] {
			if diffFound {
				return false
			}
			diffFound = true
		}
	}
	return diffFound
}
