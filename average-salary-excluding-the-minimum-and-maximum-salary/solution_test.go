package averagesalaryexcludingtheminimumandmaximumsalary

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		salary []int
		want   float64
	}{
		{[]int{4000, 3000, 1000, 2000}, 2500.00000},
		{[]int{1000, 2000, 3000}, 2000.00000},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := average(
				test.salary,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
