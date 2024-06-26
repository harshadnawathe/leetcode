package revealcardsinincreasingorder

import (
	"reflect"
	"testing"
)

func Test_deckRevealedIncreasing(t *testing.T) {
	type args struct {
		deck []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				deck: []int{17, 13, 11, 2, 3, 5, 7},
			},
			want: []int{2, 13, 3, 11, 5, 17, 7},
		},
		{
			name: "Example 2",
			args: args{
				deck: []int{1, 1000},
			},
			want: []int{1, 1000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deckRevealedIncreasing(tt.args.deck); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deckRevealedIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
