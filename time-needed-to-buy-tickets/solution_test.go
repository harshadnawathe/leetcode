package timeneededtobuytickets

import "testing"

func Test_timeRequiredToBuy(t *testing.T) {
	type args struct {
		tickets []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				tickets: []int{2, 3, 2},
				k:       2,
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				tickets: []int{5, 1, 1, 1},
				k:       0,
			},
			want: 8,
		},
		{
			name: "single person in queue",
			args: args{
				tickets: []int{2},
				k:       0,
			},
			want: 2,
		},
		{
			name: "person before needs more tickets",
			args: args{
				tickets: []int{3, 2},
				k:       1,
			},
			want: 4,
		},
		{
			name: "person after needs more tickets",
			args: args{
				tickets: []int{2, 3},
				k:       0,
			},
			want: 3,
		},
		{
			name: "person before needs less tickets",
			args: args{
				tickets: []int{3, 5},
				k:       1,
			},
			want: 8,
		},
		{
			name: "person after needs less tickets",
			args: args{
				tickets: []int{3, 1},
				k:       0,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeRequiredToBuy(tt.args.tickets, tt.args.k); got != tt.want {
				t.Errorf("timeRequiredToBuy() = %v, want %v", got, tt.want)
			}
		})
	}
}
