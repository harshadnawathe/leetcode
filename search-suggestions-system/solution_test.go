package searchsuggestionssystem

import (
	"reflect"
	"testing"
)

func Test_suggestedProducts(t *testing.T) {
	type args struct {
		products   []string
		searchWord string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"Match single word", args{[]string{"foo", "bar"}, "f"}, [][]string{{"foo"}}},
		{"Match multiple words", args{[]string{"foo", "far"}, "f"}, [][]string{{"far", "foo"}}},
		{"Match none", args{[]string{"foo", "bar"}, "x"}, [][]string{{}}},
		{
			"Example 1",
			args{
				[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse",
			},
			[][]string{
				{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"},
			},
		},
		{"Example 2", args{[]string{"havana"}, "havana"}, [][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}}},
		{"Example 3", args{[]string{"havana"}, "tatiana"}, [][]string{{}, {}, {}, {}, {}, {}, {}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := suggestedProducts(tt.args.products, tt.args.searchWord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("suggestedProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}
