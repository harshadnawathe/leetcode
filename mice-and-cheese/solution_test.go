package miceandcheese

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMiceAndCheese(t *testing.T) {
	tests := []struct {
		reward1 []int
		reward2 []int
		k       int
		want    int
	}{
		{[]int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2, 15},
		{[]int{1, 1}, []int{1, 1}, 2, 2},
		{[]int{3, 3}, []int{3, 5}, 1, 8},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := miceAndCheese(
				test.reward1,
				test.reward2,
				test.k,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
