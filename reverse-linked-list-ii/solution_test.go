package reverselinkedlistii

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				head:  listOf(1, 2, 3, 4, 5),
				left:  2,
				right: 4,
			},
			want: listOf(1, 4, 3, 2, 5),
		},
		{
			name: "Example 2",
			args: args{
				head:  listOf(5),
				left:  1,
				right: 1,
			},
			want: listOf(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
