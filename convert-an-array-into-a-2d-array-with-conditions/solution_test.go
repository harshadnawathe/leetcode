package convertanarrayintoa2darraywithconditions

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findMatrix(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{nums: []int{1, 3, 4, 1, 2, 3, 1}}, [][]int{{1, 3, 4, 2}, {1, 3}, {1}}},
		{"Example 2", args{nums: []int{1, 2, 3, 4}}, [][]int{{1, 2, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMatrix(tt.args.nums)
			sort2dArray(got)
			sort2dArray(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sort2dArray(a [][]int) {
	for i := range a {
		sort.Ints(a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		if len(a[i]) == len(a[j]) && len(a[i]) > 0 {
			return a[i][0] < a[j][0]
		}
		return len(a[i]) < len(a[j])
	})
}
