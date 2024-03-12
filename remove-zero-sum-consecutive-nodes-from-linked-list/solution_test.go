package removezerosumconsecutivenodesfromlinkedlist

import (
	"reflect"
	"testing"
)

func Test_removeZeroSumSublists(t *testing.T) {
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
				head: listOf(1, 2, -3, 3, 1),
			},
			want: listOf(3, 1),
		},
		{
			name: "Example 2",
			args: args{
				head: listOf(1, 2, 3, -3, 4),
			},
			want: listOf(1, 2, 4),
		},
		{
			name: "Example 3",
			args: args{
				head: listOf(1, 2, 3, -3, -2),
			},
			want: listOf(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeZeroSumSublists(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeZeroSumSublists() = %v, want %v", got, tt.want)
			}
		})
	}
}
