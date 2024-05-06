package removenodesfromlinkedlist

import (
	"reflect"
	"testing"
)

func Test_removeNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				head: listOf(5, 2, 13, 3, 8),
			},
			want: listOf(13, 8),
		},
		{
			name: "Example 2",
			args: args{
				head: listOf(1, 1, 1, 1),
			},
			want: listOf(1, 1, 1, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
