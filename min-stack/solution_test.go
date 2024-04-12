package minstack

import (
	"reflect"
	"testing"
)

type cmdFunc func(stk *MinStack, args ...any) any

var cmds = map[string]cmdFunc{
	"MinStack": func(stk *MinStack, args ...any) any {
		*stk = Constructor()
		return nil
	},
	"push": func(stk *MinStack, args ...any) any {
		stk.Push(args[0].(int))
		return nil
	},
	"pop": func(stk *MinStack, args ...any) any {
		stk.Pop()
		return nil
	},
	"top": func(stk *MinStack, args ...any) any {
		return stk.Top()
	},
	"getMin": func(stk *MinStack, args ...any) any {
		return stk.GetMin()
	},
}

func Test_MinStack(t *testing.T) {
	type args struct {
		cmds    []string
		cmdArgs [][]any
	}

	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "Example 1",
			args: args{
				cmds:    []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
				cmdArgs: [][]any{{}, {-2}, {0}, {-3}, {}, {}, {}, {}},
			},
			want: []any{nil, nil, nil, nil, -3, nil, 0, -2},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := []any{}
			var set MinStack
			for i, cmd := range test.args.cmds {
				got = append(got, cmds[cmd](&set, test.args.cmdArgs[i]...))
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
