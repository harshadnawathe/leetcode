package reorganizestring

import "testing"

func Test_reorganizeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				s: "aab",
			},
			want: "aba",
		},
		{
			name: "Example 2",
			args: args{
				s: "aaab",
			},
			want: "",
		},
		{
			name: "Example 3",
			args: args{
				s: "abcdefghijklmnopqrstuvwxyz",
			},
			want: "anbocpdqerfsgthuivjwkxlymz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorganizeString(tt.args.s); got != tt.want {
				t.Errorf("reorganizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
