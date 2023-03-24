package integertoroman

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{3, "III"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {

			got := intToRoman(test.num)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
