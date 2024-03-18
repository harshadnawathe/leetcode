package minimumnumberofarrowstoburstballoons

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				points: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{
				points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			},
			want: 2,
		},
		{
			name: "Overlapping in ballons",
			args: args{
				points: [][]int{{1, 7}, {2, 4}, {5, 6}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
