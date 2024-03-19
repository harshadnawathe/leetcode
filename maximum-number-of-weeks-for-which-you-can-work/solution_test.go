package maximumnumberofweeksforwhichyoucanwork

import "testing"

func Test_numberOfWeeks(t *testing.T) {
	type args struct {
		milestones []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				milestones: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				milestones: []int{5, 2, 1},
			},
			want: 7,
		},
		{
			name: "Example 3",
			args: args{
				milestones: []int{5, 9, 4, 4, 8, 9, 9, 8, 7, 3},
			},
			want: 66,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWeeks(tt.args.milestones); got != tt.want {
				t.Errorf("numberOfWeeks() = %v, want %v", got, tt.want)
			}
		})
	}
}
