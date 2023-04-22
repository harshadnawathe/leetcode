package profitableschemes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProfitableSchemes(t *testing.T) {
	tests := []struct {
		n         int
		minProfit int
		group     []int
		profit    []int
		want      int
	}{
		{5, 3, []int{2, 2}, []int{2, 3}, 2},
		{10, 5, []int{2, 3, 5}, []int{6, 7, 8}, 7},
		{1, 1, []int{2, 2, 2, 2, 2}, []int{1, 2, 1, 1, 0}, 0},
		{64, 0, []int{80, 40}, []int{88, 88}, 2},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := profitableSchemes(
				test.n,
				test.minProfit,
				test.group,
				test.profit,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
