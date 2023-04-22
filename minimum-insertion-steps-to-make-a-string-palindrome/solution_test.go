package minimuminsertionstepstomakeastringpalindrome

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinInsertions(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"zzazz", 0},
		{"mbadm", 2},
		{"leetcode", 5},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := minInsertions(
				test.s,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
