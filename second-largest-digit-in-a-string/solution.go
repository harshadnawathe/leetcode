package secondlargestdigitinastring

func secondHighest(s string) int {
	highest, secHighest := -1, -1

	for i := 0; i < len(s); i++ {
		if c := s[i]; c >= '0' && c <= '9' {
			if d := int(c - '0'); highest < d {
				secHighest, highest = highest, d
			} else if secHighest < d && d < highest {
				secHighest = d
			}
		}
	}

	return secHighest
}
