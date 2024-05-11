package minimumcosttohirekworkers

import (
	"math"
	"testing"
)

func Test_mincostToHireWorkers(t *testing.T) {
	type args struct {
		quality []int
		wage    []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Example 1",
			args: args{
				quality: []int{10, 20, 5},
				wage:    []int{70, 50, 30},
				k:       2,
			},
			want: 105.0,
		},
		{
			name: "Example 2",
			args: args{
				quality: []int{3, 1, 10, 10, 1},
				wage:    []int{4, 8, 2, 2, 7},
				k:       3,
			},
			want: 30.66667,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mincostToHireWorkers(tt.args.quality, tt.args.wage, tt.args.k)

			if roundFloat(got, 5) != tt.want {
				t.Errorf("mincostToHireWorkers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
