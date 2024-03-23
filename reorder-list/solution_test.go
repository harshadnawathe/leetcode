package reorderlist

import "testing"

func Test_reorderList(t *testing.T) {
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
				head: listOf(1, 2, 3, 4),
			},
			want: listOf(1, 4, 2, 3),
		},
		{
			name: "Example 2",
			args: args{
				head: listOf(1, 2, 3, 4, 5),
			},
			want: listOf(1, 5, 2, 4, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)

			if !equal(tt.args.head, tt.want) {
				t.Errorf("reorderList() = %v, want %v", tt.args.head, tt.want)
			}
		})
	}
}
