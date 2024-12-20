package reverseoddlevelsofbinarytree

import (
	"reflect"
	"testing"
)

func Test_reverseOddLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root: treeOf(2, 3, 5, 8, 13, 21, 34),
			},
			want: treeOf(2, 5, 3, 8, 13, 21, 34),
		},
		{
			name: "Example 2",
			args: args{
				root: treeOf(7, 13, 11),
			},
			want: treeOf(7, 11, 13),
		},
		{
			name: "Example 3",
			args: args{
				root: treeOf(0, 1, 2, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2),
			},
			want: treeOf(0, 2, 1, 0, 0, 0, 0, 2, 2, 2, 2, 1, 1, 1, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOddLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseOddLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
