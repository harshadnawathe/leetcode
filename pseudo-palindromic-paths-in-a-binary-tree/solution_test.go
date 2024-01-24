package pseudopalindromicpathsinabinarytree

import "testing"

func Test_pseudoPalindromicPaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{treeOf(2, 3, 1, 3, 1, nil, 1)}, 2},
		{"Example 2", args{treeOf(2, 1, 1, 1, 3, nil, nil, nil, nil, nil, 1)}, 1},
		{"Example 3", args{treeOf(9)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pseudoPalindromicPaths(tt.args.root); got != tt.want {
				t.Errorf("pseudoPalindromicPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
