package diameterofbinarytree

import "testing"

func Test_diameterOfBinaryTree(t *testing.T) {
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
			args: args{treeOf(1, 2, 3, 4, 5)},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{treeOf(1, 2)},
			want: 1,
		},
		{
			name: "Single node tree",
			args: args{treeOf(1)},
			want: 0,
		},
		{
			name: "Empty tree",
			args: args{nil},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
