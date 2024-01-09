package leafsimilartrees

import "testing"

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{treeOf(3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4), treeOf(3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8)}, true},
		{"Example 2", args{treeOf(1, 2, 3), treeOf(1, 3, 2)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
