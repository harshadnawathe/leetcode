package destinationcity

func destCity(paths [][]string) string {
	sourceCities := make(map[string]struct{})
	for _, path := range paths {
		sourceCities[path[0]] = struct{}{}
	}

	for _, path := range paths {
		dest := path[1]
		if _, found := sourceCities[dest]; !found {
			return dest
		}
	}

	return ""
}
