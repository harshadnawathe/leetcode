package kthlargestsuminabinarytree

import "testing"

func Test_kthLargestLevelSum(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				root: treeOf(5, 8, 9, 2, 1, 3, 7, 4, 6),
				k:    2,
			},
			want: 13,
		},
		{
			name: "Example 2",
			args: args{
				root: treeOf(1, 2, nil, 3),
				k:    1,
			},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{
				root: treeOf(1, 2, 3),
				k:    3,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestLevelSum(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargestLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
