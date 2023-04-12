package simplifypath

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSimplifyPath(t *testing.T) {

	tests := []struct {
		path string
		want string
	}{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo", "/home/foo"},
		{"/a/./b", "/a/b"},
		{"/a/.../b", "/a/.../b"},
		{"/a/x/../b", "/a/b"},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			got := simplifyPath(
				test.path,
			)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}
