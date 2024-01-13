package minimumnumberofstepstomaketwostringsanagram

import "testing"

func Test_minSteps(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"bab", "aba"}, 1},
		{"Example 2", args{"leetcode", "practice"}, 5},
		{"Example 3", args{"anagram", "mangaar"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
