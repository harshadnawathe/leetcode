package removesubfoldersfromthefilesystem

import (
	"reflect"
	"sort"
	"testing"
)

func Test_removeSubfolders(t *testing.T) {
	type args struct {
		folders []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				folders: []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"},
			},
			want: []string{"/a", "/c/d", "/c/f"},
		},
		{
			name: "Example 2",
			args: args{
				folders: []string{"/a", "/a/b/c", "/a/b/d"},
			},
			want: []string{"/a"},
		},
		{
			name: "Example 3",
			args: args{
				folders: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
			},
			want: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeSubfolders(tt.args.folders)
			sort.Strings(tt.want)
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeSubfolders() = %v, want %v", got, tt.want)
			}
		})
	}
}
