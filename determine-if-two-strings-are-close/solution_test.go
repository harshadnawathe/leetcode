package determineiftwostringsareclose

import "testing"

func Test_closeStrings(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{"abc", "bca"}, true},
		{"Example 2", args{"a", "aa"}, false},
		{"Example 3", args{"cabbba", "abbccc"}, true},
		{"Example 4", args{"abbzzca", "babzzcz"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closeStrings(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("closeStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
