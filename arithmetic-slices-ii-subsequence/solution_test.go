package arithmeticslicesiisubsequence

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"nums length is 2", args{[]int{1, 2}}, 0},
		{"nums length is 3 and diff is same", args{[]int{1, 2, 3}}, 1},
		{"nums length is 3 and diff is not same", args{[]int{1, 2, 4}}, 0},
		{"Example 1", args{[]int{2, 4, 6, 8, 10}}, 7},
		{"Example 2", args{[]int{7, 7, 7, 7, 7}}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlices(tt.args.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
