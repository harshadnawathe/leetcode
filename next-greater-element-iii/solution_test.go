package nextgreaterelementiii

import "testing"

func Test_nextGreaterElement(t *testing.T) {
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
			args: args{
				n: 12,
			},
			want: 21,
		},
		{
			name: "Example 2",
			args: args{
				n: 21,
			},
			want: -1,
		},
		{
			name: "Example 3",
			args: args{
				n: 11,
			},
			want: -1,
		},
		{
			name: "Example 4",
			args: args{
				n: 2147483486,
			},
			want: -1,
		},
		{
			name: "Example 5",
			args: args{
				n: 2147483476,
			},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.n); got != tt.want {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
