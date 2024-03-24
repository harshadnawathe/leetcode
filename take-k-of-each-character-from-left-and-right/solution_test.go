package takekofeachcharacterfromleftandright

import "testing"

func Test_takeCharacters(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				s: "aabaaaacaabc",
				k: 2,
			},
			want: 8,
		},
		{
			name: "Example 2",
			args: args{
				s: "a",
				k: 1,
			},
			want: -1,
		},
		{
			name: "Example 3",
			args: args{
				s: "cbbac",
				k: 1,
			},
			want: 3,
		},
		{
			name: "k = 0",
			args: args{
				s: "a",
				k: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := takeCharacters(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("takeCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
