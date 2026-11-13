package processstringwithspecialoperationsi

import (
	"reflect"
	"testing"
)

func Test_processStr(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "a#b%*",
			want: "ba",
		},
		{
			name: "Example 2",
			s: "z*#",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := processStr(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
