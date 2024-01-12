package amountoftimeforbinarytreetobeinfected

import "testing"

func Test_amountOfTime(t *testing.T) {
	type args struct {
		root  *TreeNode
		start int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{treeOf(1, 5, 3, nil, 4, 10, 6, 9, 2), 3}, 4},
		{"Example 2", args{treeOf(1), 1}, 0},
		{"Example 3", args{treeOf(1, 2, nil, 3, nil, 4, nil, 5), 5}, 4},
		{"Example 4", args{treeOf(2, 5), 5}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := amountOfTime(tt.args.root, tt.args.start); got != tt.want {
				t.Errorf("amountOfTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
