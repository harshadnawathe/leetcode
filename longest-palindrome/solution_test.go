package longestpalindrome

import "testing"

func Test_longestPalindrome(t *testing.T) {
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
				s: "abccccdd",
			},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				s: "bb",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
