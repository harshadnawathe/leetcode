package customsortstring

import "strings"

// time complexity O(n)
// space complexity O(n)
func customSortString(order string, s string) string {
	frequency := make([]int, 26)
	for i := 0; i < len(s); i++ {
		frequency[s[i]-'a']++
	}

	var b strings.Builder

	for i := 0; i < len(order); i++ {
		c := order[i] - 'a'
		if frequency[c] == 0 {
			continue
		}

		for frequency[c] > 0 {
			b.WriteByte(order[i])
			frequency[c]--
		}
	}

	for c, f := range frequency {
		for i := 0; i < f; i++ {
			b.WriteByte(byte('a' + c))
		}
	}

	return b.String()
}

// // time complexity O(nlogn)
// // space complexity O(n)
// func customSortString(order string, s string) string {
// 	table := make(map[byte]int)
// 	for i := 0; i < len(order); i++ {
// 		table[order[i]] = i
// 	}
//
// 	weight := func(b byte) int {
// 		if index, ok := table[b]; ok {
// 			return index
// 		}
// 		return math.MaxInt32
// 	}
//
// 	bs := []byte(s)
// 	sort.Slice(bs, func(i, j int) bool {
// 		return weight(bs[i]) < weight(bs[j])
// 	})
//
// 	return string(bs)
// }
