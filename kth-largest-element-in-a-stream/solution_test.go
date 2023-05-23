package kthlargestelementinastream

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		cmds []string
		args [][]any
		want []any
	}{
		{
			[]string{"KthLargest", "add", "add", "add", "add", "add"},
			[][]any{{3, []int{4, 5, 8, 2}}, {3}, {5}, {10}, {9}, {4}},
			[]any{nil, 4, 5, 5, 8, 8},
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

const setKey = "kThLargest"

type cmdFunc func(map[string]any, ...any) (any, error)

var cmdFuncs = map[string]cmdFunc{
	"KthLargest": func(ctx map[string]any, args ...any) (any, error) {
		if _, ok := ctx[setKey]; ok {
			return nil, errors.New("set already exists in context")
		}
		set := Constructor(args[0].(int), args[1].([]int))
		ctx[setKey] = &set
		return nil, nil
	},
	"add": func(ctx map[string]any, args ...any) (any, error) {
		if _, ok := ctx[setKey]; !ok {
			return nil, errors.New("set not found in the context")
		}

		set := ctx[setKey].(*KthLargest)
		result := set.Add(args[0].(int))
		return result, nil
	},
}
