package rangesumofbst

import "testing"

func Test_rangeSumBST(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{treeOf(10, 5, 15, 3, 7, nil, 18), 7, 15}, 32},
		{"Example 2", args{treeOf(10, 5, 15, 3, 7, 13, 18, 1, nil, 6), 6, 10}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSumBST(tt.args.root, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("rangeSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
