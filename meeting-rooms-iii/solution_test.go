package meetingroomsiii

import "testing"

func Test_mostBooked(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		meetings [][]int
		want     int
	}{
		{
			name:     "Example 1",
			n:        2,
			meetings: [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}},
			want:     0,
		},
		{
			name:     "Example 2",
			n:        3,
			meetings: [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}},
			want:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mostBooked(tt.n, tt.meetings)
			if got != tt.want {
				t.Errorf("mostBooked() = %v, want %v", got, tt.want)
			}
		})
	}
}
