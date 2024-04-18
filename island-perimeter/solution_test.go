package islantperimeter

import "testing"

func Test_islandPerimeter(t *testing.T) {
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
					{0, 1, 0, 0},
					{1, 1, 1, 0},
					{0, 1, 0, 0},
					{1, 1, 0, 0},
				},
			},
			want: 16,
		},
		{
			name: "Example 2",
			args: args{
				grid: [][]int{{1}},
			},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{
				grid: [][]int{{1, 0}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := islandPerimeter(tt.args.grid); got != tt.want {
				t.Errorf("islandPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
