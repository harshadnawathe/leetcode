package partitionlist

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
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
				head: listOf(1, 4, 3, 2, 5, 2),
				x:    3,
			},
			want: listOf(1, 2, 2, 4, 3, 5),
		},
		{
			name: "Example 2",
			args: args{
				head: listOf(2, 1),
				x:    2,
			},
			want: listOf(1, 2),
		},
		{
			name: "Example 3",
			args: args{
				head: listOf(),
				x:    2,
			},
			want: listOf(),
		},
		{
			name: "Example 4",
			args: args{
				head: listOf(1, 2, 3),
				x:    -1,
			},
			want: listOf(1, 2, 3),
		},
		{
			name: "Example 5",
			args: args{
				head: listOf(1, 2, 3),
				x:    4,
			},
			want: listOf(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
