package minimumoperationstoformsubsequencewithtargetsum

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums:   []int{1, 2, 8},
				target: 7,
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				nums:   []int{1, 32, 1, 2},
				target: 12,
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				nums:   []int{1, 32, 1},
				target: 35,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
