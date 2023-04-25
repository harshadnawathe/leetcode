package smallestnumberininfiniteset

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		cmds []string
		args [][]int
		want []any
	}{
		{
			[]string{"SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"},
			[][]int{{}, {2}, {}, {}, {}, {1}, {}, {}, {}},
			[]any{nil, nil, 1, 2, 3, nil, 1, 4, 5},
		},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Test_%v", i)
		t.Run(testName, func(t *testing.T) {
			var got []any
			ctx := make(map[string]any)
			for j, cmd := range test.cmds {
				res, err := cmdFuncs[cmd](ctx, test.args[j]...)
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

const setKey = "smallestInfiniteSet"

type cmdFunc func(map[string]any, ...int) (any, error)

var cmdFuncs = map[string]cmdFunc{
	"SmallestInfiniteSet": func(ctx map[string]any, args ...int) (any, error) {
		if _, ok := ctx[setKey]; ok {
			return nil, errors.New("set already exists in context")
		}
		set := Constructor()
		ctx[setKey] = &set
		return nil, nil
	},
	"addBack": func(ctx map[string]any, args ...int) (any, error) {
		if _, ok := ctx[setKey]; !ok {
			return nil, errors.New("set not found in the context")
		}

		set := ctx[setKey].(*SmallestInfiniteSet)
		set.AddBack(args[0])
		return nil, nil
	},
	"popSmallest": func(ctx map[string]any, args ...int) (any, error) {
		if _, ok := ctx[setKey]; !ok {
			return nil, errors.New("set not found in the context")
		}

		set := ctx[setKey].(*SmallestInfiniteSet)
		x := set.PopSmallest()
		return x, nil
	},
}
