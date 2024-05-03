package compareversionnumbers

import (
	"strconv"
)

func compareVersion(version1 string, version2 string) int {
	return compare(
		revisions(version1),
		revisions(version2),
	)
}

func revisions(version string) []int {
	revs := []int{}

	for b := 0; b < len(version); b++ {
		e := b + 1
		for ; e < len(version) && version[e] != '.'; e++ {
		}

		n, _ := strconv.Atoi(version[b:e])
		revs = append(revs, n)

		b = e
	}

	return revs
}

func compare(rev1, rev2 []int) int {
	for i1, i2 := 0, 0; i1 < len(rev1) || i2 < len(rev2); i1, i2 = i1+1, i2+1 {
		var x1 int
		if i1 < len(rev1) {
			x1 = rev1[i1]
		}

		var x2 int
		if i2 < len(rev2) {
			x2 = rev2[i2]
		}

		d := x1 - x2
		if d < 0 {
			return -1
		} else if d > 0 {
			return 1
		}

	}

	return 0
}
