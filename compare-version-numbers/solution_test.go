package compareversionnumbers

import "testing"

func Test_compareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				version1: "1.01",
				version2: "1.001",
			},
			want: 0,
		},
		{
			name: "Example 2",
			args: args{
				version1: "1.0",
				version2: "1.0.0",
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				version1: "0.1",
				version2: "1.1",
			},
			want: -1,
		},
		{
			name: "Example 4",
			args: args{
				version1: "1.0.1",
				version2: "1",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
