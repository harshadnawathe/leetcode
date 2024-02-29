package evenoddtree

import "testing"

func Test_isEvenOddTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				root: treeOf(1, 10, 4, 3, nil, 7, 9, 12, 8, 6, nil, nil, 2),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				root: treeOf(5, 4, 2, 3, 3, 7),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEvenOddTree(tt.args.root); got != tt.want {
				t.Errorf("isEvenOddTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
