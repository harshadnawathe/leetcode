package pascalstriangleii

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		rowIndex int
		want     []int
	}{
		{0, []int{1}},
		{1, []int{1, 1}},
		{3, []int{1, 3, 3, 1}},
		{4, []int{1, 4, 6, 4, 1}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := getRow(
				test.rowIndex,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
