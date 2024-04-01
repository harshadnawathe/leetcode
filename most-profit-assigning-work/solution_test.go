package mostprofitassigningwork

import "testing"

func Test_maxProfitAssignment(t *testing.T) {
	type args struct {
		difficulty []int
		profit     []int
		worker     []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				difficulty: []int{2, 4, 6, 8, 10},
				profit:     []int{10, 20, 30, 40, 50},
				worker:     []int{4, 5, 6, 7},
			},
			want: 100,
		},
		{
			name: "Example 2",
			args: args{
				difficulty: []int{85, 47, 57},
				profit:     []int{24, 66, 99},
				worker:     []int{40, 25, 25},
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				difficulty: []int{68, 35, 52, 47, 86},
				profit:     []int{67, 17, 1, 81, 3},
				worker:     []int{92, 10, 85, 84, 82},
			},
			want: 324,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitAssignment(tt.args.difficulty, tt.args.profit, tt.args.worker); got != tt.want {
				t.Errorf("maxProfitAssignment() = %v, want %v", got, tt.want)
			}
		})
	}
}
