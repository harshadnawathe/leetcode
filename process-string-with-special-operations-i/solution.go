package processstringwithspecialoperationsi

func processStr(s string) string {
	var result []rune
	for _, ch := range s {
		switch ch {
		case '*':
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		case '#':
			dup := make([]rune, len(result), 2*len(result))
			copy(dup, result)
			result = append(dup, result...)
		case '%':
			for b, e := 0, len(result)-1; b < e; b, e = b+1, e-1 {
				result[b], result[e] = result[e], result[b]
			}
		default:
			result = append(result, ch)
		}
	}

	return string(result)
}
