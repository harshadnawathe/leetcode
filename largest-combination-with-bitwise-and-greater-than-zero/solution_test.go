package largestcombinationwithbitwiseandgreaterthanzero

import "testing"

func Test_largestCombination(t *testing.T) {
	type args struct {
		candidates []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				candidates: []int{16, 17, 71, 62, 12, 24, 14},
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				candidates: []int{8, 8},
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				candidates: []int{84, 40, 66, 44, 91, 90, 1, 14, 73, 51, 47, 35, 18, 46, 18, 65, 55, 18, 16, 45, 43, 58, 90, 92, 91, 43, 44, 76, 85, 72, 24, 89, 60, 94, 81, 90, 86, 79, 84, 41, 41, 28, 44},
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestCombination(tt.args.candidates); got != tt.want {
				t.Errorf("largestCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}
