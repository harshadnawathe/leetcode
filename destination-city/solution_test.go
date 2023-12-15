package destinationcity

import "testing"

func Test_destCity(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				paths: [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
			},
			want: "Sao Paulo",
		},
		{
			name: "Example 2",
			args: args{
				paths: [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}},
			},
			want: "A",
		},
		{
			name: "Example 3",
			args: args{
				paths: [][]string{{"A", "Z"}},
			},
			want: "Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := destCity(tt.args.paths); got != tt.want {
				t.Errorf("destCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
