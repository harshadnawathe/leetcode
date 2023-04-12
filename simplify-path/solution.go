package simplifypath

import (
	"strings"
)

func simplifyPath(path string) string {
	var segments []string
	for start, i := 1, 1; i < len(path); i++ {
		for i < len(path) && path[i] != '/' {
			i++
		}
		switch segment := path[start:i]; segment {
		case "":
		case ".":
		case "..":
			if len(segments) > 0 {
				segments = segments[:len(segments)-1]
			}
		default:
			segments = append(segments, segment)
		}

		start = i + 1
	}

	return "/" + strings.Join(segments, "/")
}
