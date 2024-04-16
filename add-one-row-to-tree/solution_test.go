package addonerowtotree

import (
	"reflect"
	"testing"
)

func Test_addOneRow(t *testing.T) {
	type args struct {
		root  *TreeNode
		val   int
		depth int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root:  treeOf(4, 2, 6, 3, 1, 5),
				val:   1,
				depth: 2,
			},
			want: treeOf(4, 1, 1, 2, nil, nil, 6, 3, 1, 5),
		},
		{
			name: "Example 2",
			args: args{
				root:  treeOf(4, 2, nil, 3, 1),
				val:   1,
				depth: 3,
			},
			want: treeOf(4, 2, nil, 1, 1, 3, nil, nil, 1),
		},
		{
			name: "insert row at depth 1",
			args: args{
				root:  treeOf(1, 2, 3),
				val:   99,
				depth: 1,
			},
			want: treeOf(99, 1, nil, 2, 3),
		},
		{
			name: "insert row at the bottom",
			args: args{
				root:  treeOf(1, 2, 3),
				val:   99,
				depth: 3,
			},
			want: treeOf(1, 2, 3, 99, 99, 99, 99),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOneRow(tt.args.root, tt.args.val, tt.args.depth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOneRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
