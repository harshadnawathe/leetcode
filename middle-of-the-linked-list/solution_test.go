package middleofthelinkedlist

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "List with single node",
			args: args{
				head: listOf(1),
			},
			want: listOf(1),
		},
		{
			name: "List with two nodes",
			args: args{
				head: listOf(1, 2),
			},
			want: listOf(2),
		},
		{
			name: "List with three nodes",
			args: args{
				head: listOf(1, 2, 3),
			},
			want: listOf(2, 3),
		},
		{
			name: "Example 1",
			args: args{
				head: listOf(1, 2, 3, 4, 5),
			},
			want: listOf(3, 4, 5),
		},
		{
			name: "Example 2",
			args: args{
				head: listOf(1, 2, 3, 4, 5, 6),
			},
			want: listOf(4, 5, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
