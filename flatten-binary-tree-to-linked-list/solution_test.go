package flattenbinarytreetolinkedlist

import (
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
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
				root: treeOf(1, 2, 5, 3, 4, nil, 6),
			},
			want: treeOf(1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6),
		},
		{
			name: "Example 2",
			args: args{
				root: treeOf(),
			},
			want: treeOf(),
		},
		{
			name: "Example 3",
			args: args{
				root: treeOf(0),
			},
			want: treeOf(0),
		},
		{
			name: "tree with only left child",
			args: args{
				root: treeOf(2, 1),
			},
			want: treeOf(2, nil, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.root)
			if !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("flatten() = %v, want %v", tt.args.root, tt.want)
			}
		})
	}
}
