package maximalrectangle

import "testing"

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				matrix: [][]byte{
					{'1', '0', '1', '0', '0'},
					{'1', '0', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '0', '0', '1', '0'},
				},
			},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{
				matrix: [][]byte{{'0'}},
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				matrix: [][]byte{{'1'}},
			},
			want: 1,
		},
		{
			name: "Example 4",
			args: args{
				matrix: [][]byte{{'0', '1'}, {'1', '0'}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
