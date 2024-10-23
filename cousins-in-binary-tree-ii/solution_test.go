package cousinsinbinarytreeii

import (
	"reflect"
	"testing"
)

func Test_replaceValueInTree(t *testing.T) {
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
				root: treeOf(5, 4, 9, 1, 10, nil, 7),
			},
			want: treeOf(0, 0, 0, 7, 7, nil, 11),
		},
		{
			name: "Example 2",
			args: args{
				root: treeOf(3, 1, 2),
			},
			want: treeOf(0, 0, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceValueInTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceValueInTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
