package longestcommonprefix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {

			got := longestCommonPrefix(test.strs)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
