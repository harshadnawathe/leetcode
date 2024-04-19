package minimumgeneticmutation

import "testing"

func Test_minMutation(t *testing.T) {
	type args struct {
		startGene string
		endGene   string
		bank      []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				startGene: "AACCGGTT",
				endGene:   "AACCGGTA",
				bank:      []string{"AACCGGTA"},
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				startGene: "AACCGGTT",
				endGene:   "AAACGGTA",
				bank:      []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutation(tt.args.startGene, tt.args.endGene, tt.args.bank); got != tt.want {
				t.Errorf("minMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
