package constructstringwithrepeatlimit

import "testing"

func Test_repeatLimitedString(t *testing.T) {
	type args struct {
		s           string
		repeatLimit int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				s:           "cczazcc",
				repeatLimit: 3,
			},
			want: "zzcccac",
		},
		{
			name: "Example 2",
			args: args{
				s:           "aababab",
				repeatLimit: 2,
			},
			want: "bbabaa",
		},
		{
			name: "Example 3",
			args: args{
				s:           "a",
				repeatLimit: 1,
			},
			want: "a",
		},
		{
			name: "Example 4",
			args: args{
				s:           "aabbcc",
				repeatLimit: 1,
			},
			want: "cbcba",
		},
		{
			name: "Example 5",
			args: args{
				s:           "aaaaa",
				repeatLimit: 2,
			},
			want: "aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatLimitedString(tt.args.s, tt.args.repeatLimit); got != tt.want {
				t.Errorf("repeatLimitedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
