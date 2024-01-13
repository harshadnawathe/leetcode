package minimumnumberofstepstomaketwostringsanagram

// time: O(N)
// space: O(1)
func minSteps(s string, t string) int {
	table := charFreqTable(s)
	for i := 0; i < len(t); i++ {
		if table[t[i]-'a'] > 0 {
			table[t[i]-'a']--
		}
	}
	steps := 0
	for _, n := range table {
		steps += n
	}
	return steps
}

func charFreqTable(s string) []int {
	table := make([]int, 26)
	for i := 0; i < len(s); i++ {
		table[s[i]-'a']++
	}
	return table
}
