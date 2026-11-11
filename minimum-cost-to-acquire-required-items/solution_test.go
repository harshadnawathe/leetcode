package minimumcosttoacquirerequireditems

import "testing"

func Test_minimumCost(t *testing.T) {
	tests := []struct {
		name     string
		cost1    int
		cost2    int
		costBoth int
		need1    int
		need2    int
		want     int64
	}{
		{
			name:     "Example 1",
			cost1:    3,
			cost2:    2,
			costBoth: 1,
			need1:    3,
			need2:    2,
			want:     3,
		},
		{
			name:     "Example 2",
			cost1:    5,
			cost2:    4,
			costBoth: 15,
			need1:    2,
			need2:    3,
			want:     22,
		},
		{
			name:     "Example 3",
			cost1:    5,
			cost2:    4,
			costBoth: 15,
			need1:    0,
			need2:    0,
			want:     0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumCost(tt.cost1, tt.cost2, tt.costBoth, tt.need1, tt.need2)
			if got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
