package minimumcostfortickets

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMincostTickets(t *testing.T) {

	tests := []struct {
		days  []int
		costs []int
		want  int
	}{
		//{[]int{}, []int{2, 7, 15}, 0},
		{[]int{1}, []int{2, 7, 15}, 2},
		{[]int{1, 4}, []int{2, 7, 15}, 4},
		{[]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}, 11},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}, 17},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := mincostTickets(
				test.days,
				test.costs,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
