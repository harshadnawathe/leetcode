package pattern132

import (
	"fmt"
	"testing"
)

func Test_find132pattern(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{[]int{1, 2, 3, 4}}, false},
		{args{[]int{3, 1, 4, 2}}, true},
		{args{[]int{-1, 3, 2, 0}}, true},
		{args{[]int{1}}, false},
		{args{[]int{1, 2}}, false},
		{args{[]int{1, 3, 2}}, true},
		{args{[]int{3, 5, 0, 3, 4}}, true},
	}
	for i, tt := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			if got := find132pattern(tt.args.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
