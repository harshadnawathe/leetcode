package deletenodeinalinkedlist

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		vals []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				vals: []int{4, 5, 1, 9},
				x:    5,
			},
			want: listOf(4, 1, 9),
		},
		{
			name: "Example 2",
			args: args{
				vals: []int{4, 5, 1, 9},
				x:    1,
			},
			want: listOf(4, 5, 9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := listOf(tt.args.vals...)
			deleteNode(findNode(list, tt.args.x))

			if !reflect.DeepEqual(list, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", list, tt.want)
			}
		})
	}
}
