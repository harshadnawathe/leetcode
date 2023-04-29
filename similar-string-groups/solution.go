package similarstringgroups

func numSimilarGroups(strs []string) int {
	graph := makeAdjacencyList(strs)
	visited := make(map[string]struct{})
	var dfs func(string)
	dfs = func(s string) {
		if _, found := visited[s]; found {
			return
		}
		visited[s] = struct{}{}

		for _, next := range graph[s] {
			dfs(next)
		}
	}

	count := 0
	for _, str := range strs {
		if _, found := visited[str]; found {
			continue
		}
		dfs(str)
		count++
	}

	return count
}

func isSimilar(s1, s2 string) bool {
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff <= 2
}

func makeAdjacencyList(strs []string) map[string][]string {
	graph := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if isSimilar(strs[i], strs[j]) {
				graph[strs[i]] = append(graph[strs[i]], strs[j])
				graph[strs[j]] = append(graph[strs[j]], strs[i])
			}
		}
	}
	return graph
}
