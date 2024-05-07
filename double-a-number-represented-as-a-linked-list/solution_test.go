package doubleanumberrepresentedasalinkedlist

import (
	"reflect"
	"testing"
)

func Test_doubleIt(t *testing.T) {
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
				head: listOf(1, 8, 9),
			},
			want: listOf(3, 7, 8),
		},
		{
			name: "Example 2",
			args: args{
				head: listOf(9, 9, 9),
			},
			want: listOf(1, 9, 9, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doubleIt(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doubleIt() = %v, want %v", got, tt.want)
			}
		})
	}
}
