package successfulpairsofspellsandpotions

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSuccessfulPairs(t *testing.T) {

	tests := []struct {
		spells  []int
		potions []int
		success int64
		want    []int
	}{
		{[]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7, []int{4, 0, 3}},
		{[]int{3, 1, 2}, []int{8, 5, 8}, 16, []int{2, 0, 2}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := successfulPairs(
				test.spells,
				test.potions,
				test.success,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
