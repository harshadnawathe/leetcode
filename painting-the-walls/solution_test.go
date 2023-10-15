package paintingthewalls

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPaintWalls(t *testing.T) {
	tests := []struct {
		cost []int
		time []int
		want int
	}{
		{
			cost: []int{1, 2, 3, 2},
			time: []int{1, 2, 3, 2},
			want: 3,
		},
		{
			cost: []int{2, 3, 4, 2},
			time: []int{1, 1, 1, 1},
			want: 4,
		},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := paintWalls(
				test.cost,
				test.time,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
