package minimizemaximumofarray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinimizeArrayValue(t *testing.T) {

	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 7, 1, 6}, 5},
		{[]int{10, 1}, 10},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := minimizeArrayValue(
				test.nums,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
