package isomorphicstrings

func isIsomorphic(s string, t string) bool {
	sMap, tMap := make(map[byte]byte), make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		_, sOk := sMap[s[i]]
		_, tOk := tMap[t[i]]

		if !sOk && !tOk {
			sMap[s[i]] = t[i]
			tMap[t[i]] = s[i]
		} else if sOk {
			if sMap[s[i]] != t[i] {
				return false
			}
		} else if tOk {
			if tMap[t[i]] != s[i] {
				return false
			}
		}
	}

	return true
}
