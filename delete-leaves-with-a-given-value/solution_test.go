package deleteleaveswithagivenvalue

import (
	"reflect"
	"testing"
)

func Test_removeLeafNodes(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example 1",
			args: args{
				root:   treeOf(1, 2, 3, 2, nil, 2, 4),
				target: 2,
			},
			want: treeOf(1, nil, 3, nil, 4),
		},
		{
			name: "Example 2",
			args: args{
				root:   treeOf(1, 3, 3, 3, 2),
				target: 3,
			},
			want: treeOf(1, 3, nil, nil, 2),
		},
		{
			name: "Example 3",
			args: args{
				root:   treeOf(1, 2, nil, 2, nil, 2),
				target: 2,
			},
			want: treeOf(1),
		},
		{
			name: "Example 4",
			args: args{
				root:   treeOf(1),
				target: 1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeLeafNodes(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeLeafNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
