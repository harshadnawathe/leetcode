package minimumnumberofchangestomakebinarystringbeautiful

func minChanges(s string) int {
	var changeCount int
	for i := 0; i < len(s); i += 2 {
		if s[i] != s[i+1] {
			changeCount++
		}
	}
	return changeCount
}
