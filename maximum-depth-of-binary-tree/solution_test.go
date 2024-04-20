package maximumdepthofbinarytree

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				root: treeOf(3, 9, 20, nil, nil, 15, 7),
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				root: treeOf(1, nil, 2),
			},
			want: 2,
		},
		{
			name: "nil treee",
			args: args{
				root: treeOf(),
			},
			want: 0,
		},
		{
			name: "single node",
			args: args{
				root: treeOf(1),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
