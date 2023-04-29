package similarstringgroups

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumSimilarGroups(t *testing.T) {
	tests := []struct {
		strs []string
		want int
	}{
		{[]string{"tars", "rats", "arts", "star"}, 2},
		{[]string{"omv", "ovm"}, 1},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := numSimilarGroups(
				test.strs,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
