package secondlargestdigitinastring

import "testing"

func Test_secondHighest(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				s: "dfa12321afd",
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				s: "abc1111",
			},
			want: -1,
		},
		{
			name: "Example 3",
			args: args{
				s: "ck077",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondHighest(tt.args.s); got != tt.want {
				t.Errorf("secondHighest() = %v, want %v", got, tt.want)
			}
		})
	}
}
