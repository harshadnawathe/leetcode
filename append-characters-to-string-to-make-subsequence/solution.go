package appendcharacterstostringtomakesubsequence

func appendCharacters(s string, t string) int {

	ti := 0

	for si := 0; si < len(s) && ti < len(t); si++ {
		if s[si] == t[ti] {
			ti++
		}
	}

	return len(t) - ti
}
