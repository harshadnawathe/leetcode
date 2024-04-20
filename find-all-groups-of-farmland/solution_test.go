package findallgroupsoffarmland

import (
	"reflect"
	"testing"
)

func Test_findFarmland(t *testing.T) {
	type args struct {
		land [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				land: [][]int{
					{1, 0, 0},
					{0, 1, 1},
					{0, 1, 1},
				},
			},
			want: [][]int{{0, 0, 0, 0}, {1, 1, 2, 2}},
		},
		{
			name: "Example 2",
			args: args{
				land: [][]int{{1, 1}, {1, 1}},
			},
			want: [][]int{{0, 0, 1, 1}},
		},
		{
			name: "Example 3",
			args: args{
				land: [][]int{{0}},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFarmland(tt.args.land); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFarmland() = %v, want %v", got, tt.want)
			}
		})
	}
}
