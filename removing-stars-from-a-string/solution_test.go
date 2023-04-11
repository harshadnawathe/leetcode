package removingstarsfromastring

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveStars(t *testing.T) {

	tests := []struct {
		s    string
		want string
	}{
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
		{"abb*cdfg*****x*", "a"},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := removeStars(
				test.s,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
