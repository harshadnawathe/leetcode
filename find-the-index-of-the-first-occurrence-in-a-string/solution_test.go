package findtheindexofthefirstoccurrenceinastring

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStrStr(t *testing.T) {

	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"xyzabc", "abc", 3},
		{"ababcaababcaabc", "ababcaabc", 6},
		{"aaa", "aaaa", -1},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := strStr(
				test.haystack,
				test.needle,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
