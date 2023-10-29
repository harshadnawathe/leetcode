package poorpigs

import "testing"

func Test_poorPigs(t *testing.T) {
	type args struct {
		buckets       int
		minutesToDie  int
		minutesToTest int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0 Buckets", args{0, 1, 1}, 0},
		{"1 Buckets", args{1, 1, 1}, 0},
		{"2 Buckets, single test", args{2, 1, 1}, 1},
		{"2 Buckets, two tests", args{2, 1, 2}, 1},
		{"3 Buckets, single test", args{3, 1, 1}, 2},
		{"3 Buckets, two tests", args{3, 1, 2}, 1},
		{"4 Buckets, single test", args{4, 1, 1}, 2},
		{"4 Buckets, two tests", args{4, 1, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := poorPigs(tt.args.buckets, tt.args.minutesToDie, tt.args.minutesToTest); got != tt.want {
				t.Errorf("poorPigs() = %v, want %v", got, tt.want)
			}
		})
	}
}
