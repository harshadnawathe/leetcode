package sumofleftleaves

import "testing"

func Test_sumOfLeftLeaves(t *testing.T) {
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
			want: 24,
		},
		{
			name: "Example 2",
			args: args{
				root: treeOf(1),
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				root: treeOf(1, 2, 3, 4, 5, 6, 7, nil, nil, nil, nil, nil, nil, 8, nil),
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
