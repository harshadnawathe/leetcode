package scoreafterflippingmatrix

import "testing"

func Test_matrixScore(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				grid: [][]int{
					{0, 0, 1, 1},
					{1, 0, 1, 0},
					{1, 1, 0, 0},
				},
			},
			want: 39,
		},
		{
			name: "Example 2",
			args: args{
				grid: [][]int{{0}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixScore(tt.args.grid); got != tt.want {
				t.Errorf("matrixScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
