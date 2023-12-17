package designafoodratingsystem

import (
	"errors"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		cmds []string
		args [][]any
	}

	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "Example 1",
			args: args{
				cmds: []string{"FoodRatings", "highestRated", "highestRated", "changeRating", "highestRated", "changeRating", "highestRated"},
				args: [][]any{
					{
						[]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
						[]string{"korean", "japanese", "japanese", "greek", "japanese", "korean"},
						[]int{9, 12, 8, 15, 14, 7},
					},
					{"korean"},
					{"japanese"},
					{"sushi", 16},
					{"japanese"},
					{"ramen", 16},
					{"japanese"},
				},
			},
			want: []any{nil, "kimchi", "ramen", nil, "sushi", nil, "ramen"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got []any
			ctx := make(map[string]any)
			for i, cmd := range test.args.cmds {
				res, err := cmdFuncs[cmd](ctx, test.args.args[i]...)
				if err != nil {
					t.Errorf("Failed. Got error: %v", err)
				}
				got = append(got, res)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Failed. want= %v got= %v", test.want, got)
			}
		})
	}
}

const foodRatingsKey = "foodRatings"

type cmdFunc func(map[string]any, ...any) (any, error)

var cmdFuncs = map[string]cmdFunc{
	"FoodRatings": func(ctx map[string]any, args ...any) (any, error) {
		if _, ok := ctx[foodRatingsKey]; ok {
			return nil, errors.New("foodRatings already exists in context")
		}
		foodRatings := Constructor(
			args[0].([]string),
			args[1].([]string),
			args[2].([]int),
		)
		ctx[foodRatingsKey] = &foodRatings
		return nil, nil
	},
	"highestRated": func(ctx map[string]any, args ...any) (any, error) {
		if _, ok := ctx[foodRatingsKey]; !ok {
			return nil, errors.New("foodRatings already exists in context")
		}

		foodRatings := ctx[foodRatingsKey].(*FoodRatings)
		highestRated := foodRatings.HighestRated(args[0].(string))
		return highestRated, nil
	},
	"changeRating": func(ctx map[string]any, args ...any) (any, error) {
		if _, ok := ctx[foodRatingsKey]; !ok {
			return nil, errors.New("foodRatings already exists in context")
		}
		foodRatings := ctx[foodRatingsKey].(*FoodRatings)
		foodRatings.ChangeRating(args[0].(string), args[1].(int))
		return nil, nil
	},
}
