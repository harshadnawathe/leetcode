package countvowelspermutation

import "testing"

func Test_countVowelPermutation(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{1},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{2},
			want: 10,
		},
		{
			name: "Example 3",
			args: args{5},
			want: 68,
		},
		{
			name: "Example 4",
			args: args{1000},
			want: 89945857,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowelPermutation(tt.args.n); got != tt.want {
				t.Errorf("countVowelPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
