package smalleststringstartingfromleaf

import "testing"

func Test_smallestFromLeaf(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				root: treeOf(0, 1, 2, 3, 4, 3, 4),
			},
			want: "dba",
		},
		{
			name: "Example 2",
			args: args{
				root: treeOf(25, 1, 3, 1, 3, 0, 2),
			},
			want: "adz",
		},
		{
			name: "Example 3",
			args: args{
				root: treeOf(2, 2, 1, nil, 1, 0, nil, 0),
			},
			want: "abc",
		},
		{
			name: "Example 4",
			args: args{
				root: treeOf(4, 0, 1, 1),
			},
			want: "bae",
		},
		{
			name: "Example 5",
			args: args{
				root: treeOf(25, 1, nil, 0, 0, 1, nil, nil, nil, 0),
			},
			want: "ababz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestFromLeaf(tt.args.root); got != tt.want {
				t.Errorf("smallestFromLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
