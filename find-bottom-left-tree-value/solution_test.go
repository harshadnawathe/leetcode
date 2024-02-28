package findbottomlefttreevalue

import "testing"

func Test_findBottomLeftValue(t *testing.T) {
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
			args: args{treeOf(2, 1, 3)},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{treeOf(1, 2, 3, 4, nil, 5, 6, nil, nil, 7)},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
