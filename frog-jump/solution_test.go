package frogjump

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCanCross(t *testing.T) {
	tests := []struct {
		stones []int
		want   bool
	}{
		{[]int{0, 1, 3, 5, 6, 8, 12, 17}, true},
		{[]int{0, 1, 2, 3, 4, 8, 9, 11}, false},
		{[]int{0, 2}, false},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := canCross(
				test.stones,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
