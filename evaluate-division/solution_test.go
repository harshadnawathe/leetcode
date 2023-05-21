package evaluatedivision

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		{[][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}, []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}},
		{[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}, []float64{3.75000, 0.40000, 5.00000, 0.20000}},
		{[][]string{{"a", "b"}}, []float64{0.5}, [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}, []float64{0.50000, 2.00000, -1.00000, -1.00000}},
	}
	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := calcEquation(
				test.equations,
				test.values,
				test.queries,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
