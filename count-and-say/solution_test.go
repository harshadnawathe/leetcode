package countandsay

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := countAndSay(
				test.n,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
