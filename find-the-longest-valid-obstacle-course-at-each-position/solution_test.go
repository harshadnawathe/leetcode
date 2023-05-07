package findthelongestvalidobstaclecourseateachposition

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestObstacleCourseAtEachPosition(t *testing.T) {
	tests := []struct {
		obstacles []int
		want      []int
	}{
		{[]int{1, 2, 3, 2}, []int{1, 2, 3, 3}},
		{[]int{2, 2, 1}, []int{1, 2, 1}},
		{[]int{3, 1, 5, 6, 4, 2}, []int{1, 1, 2, 3, 2, 2}},
		{[]int{5, 1, 5, 5, 1, 3, 4, 5, 1, 4}, []int{1, 1, 2, 3, 2, 3, 4, 5, 3, 5}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := longestObstacleCourseAtEachPosition(
				test.obstacles,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
