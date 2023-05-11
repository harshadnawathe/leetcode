package latesttimebyreplacinghiddendigits

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaximumTime(t *testing.T) {
	tests := []struct {
		time string
		want string
	}{
		{"2?:?0", "23:50"},
		{"0?:3?", "09:39"},
		{"1?:22", "19:22"},
		{"?4:03", "14:03"},
		{"??:3?", "23:39"},
		{"??:??", "23:59"},
		{"?0:15", "20:15"},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := maximumTime(
				test.time,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
