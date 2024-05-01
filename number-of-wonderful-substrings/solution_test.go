package numberofwonderfulsubstrings

import "testing"

func Test_wonderfulSubstrings(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				word: "aba",
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				word: "aabb",
			},
			want: 9,
		},
		{
			name: "Example 3",
			args: args{
				word: "he",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wonderfulSubstrings(tt.args.word); got != tt.want {
				t.Errorf("wonderfulSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
