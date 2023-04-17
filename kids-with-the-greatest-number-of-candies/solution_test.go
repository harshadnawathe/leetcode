package kidswiththegreatestnumberofcandies

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		candies      []int
		extraCandies int
		want         []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := kidsWithCandies(
				test.candies,
				test.extraCandies,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
