package appendcharacterstostringtomakesubsequence

import "testing"

func Test_appendCharacters(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				s: "coaching",
				t: "coding",
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				s: "abcde",
				t: "a",
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				s: "z",
				t: "abcde",
			},
			want: 5,
		},
		{
			name: "Example 4",
			args: args{
				s: "apbqcrdse",
				t: "pqrs",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendCharacters(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("appendCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
