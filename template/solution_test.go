package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestXxx(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {

			got := test.input

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
